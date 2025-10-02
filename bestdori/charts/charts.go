package charts

import (
	"encoding/json"
	"math"
	"slices"
	"sort"

	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// Chart 谱面
type Chart []Note

// GetChart 获取官方谱面
func GetChart(api *uniapi.UniAPI, id int, difficulty dto.ChartDifficultyName) (*Chart, error) {
	endpoint := endpoints.ChartsInfo(id, string(difficulty))
	return uniapi.Get[Chart](api, endpoint, nil)
}

// UnmarshalSlice 从 []map[string]any 中解析谱面
func UnmarshalSlice(data []map[string]any) (*Chart, error) {
	var chart Chart = make(Chart, 0, len(data))
	for _, item := range data {
		note, err := unmarshalMap(item)
		if err != nil {
			return nil, err
		}
		chart = append(chart, *note)
	}
	return &chart, nil
}

// MarshalSlice 将谱面转换为 []map[string]any
func (c *Chart) MarshalSlice() ([]map[string]any, error) {
	bytes, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	var data []map[string]any
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Chart) Len() int {
	return len(*c)
}

func (c *Chart) Less(i, j int) bool {
	var beatI, beatJ float64
	if (*c)[i].Beat != 0.0 {
		beatI = (*c)[i].Beat
	} else {
		if len((*c)[i].Connections) > 0 {
			beatI = (*c)[i].Connections[0].Beat
		} else {
			beatI = 0.0
		}
	}
	if (*c)[j].Beat != 0.0 {
		beatJ = (*c)[j].Beat
	} else {
		if len((*c)[j].Connections) > 0 {
			beatJ = (*c)[j].Connections[0].Beat
		} else {
			beatJ = 0.0
		}
	}
	return beatI < beatJ
}

func (c *Chart) Swap(i, j int) {
	(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
}

func (c *Chart) flatten() []Note {
	var notes []Note
	for _, note := range *c {
		if note.Type == dto.NoteTypeSlide || note.Type == dto.NoteTypeLong {
			if len(note.Connections) > 0 {
				for _, conn := range note.Connections {
					notes = append(notes, Note{
						Type:   note.Type,
						Beat:   conn.Beat,
						Lane:   conn.Lane,
						Flick:  conn.Flick,
						Skill:  conn.Skill,
						Hidden: conn.Hidden,
					})
				}
			}
		} else {
			notes = append(notes, note)
		}
	}

	// 按节拍排序
	slices.SortFunc(notes, func(a, b Note) int {
		if a.Beat != 0.0 && b.Beat != 0.0 {
			if a.Beat < b.Beat {
				return -1
			} else if a.Beat > b.Beat {
				return 1
			}
		}
		return 0
	})

	// 清理后跟 BPM 音符
	for len(notes) > 0 && notes[len(notes)-1].Type == dto.NoteTypeBPM {
		notes = notes[:len(notes)-1]
	}

	return notes
}

type bpmDuration struct {
	BPM      float64
	Duration float64
}

func handleBPMDuration(duration float64, prevBPM float64, bpmDurationStack *[]bpmDuration) {
	for duration > 0.0 {
		if len(*bpmDurationStack) <= 0 {
			bpmDur := bpmDuration{
				BPM:      prevBPM,
				Duration: duration,
			}
			*bpmDurationStack = append(*bpmDurationStack, bpmDur)
			break
		} else {
			top := (*bpmDurationStack)[len(*bpmDurationStack)-1]
			if (top.BPM > 0.0 && prevBPM > 0.0) || (top.BPM < 0.0 && prevBPM < 0.0) {
				// 同向 BPM 时长不相减
				if top.BPM == prevBPM {
					top.Duration += duration
					(*bpmDurationStack)[len(*bpmDurationStack)-1] = top
				} else {
					bpmDur := bpmDuration{
						BPM:      prevBPM,
						Duration: duration,
					}
					*bpmDurationStack = append(*bpmDurationStack, bpmDur)
				}
				break
			} else {
				// 进行相减
				if top.Duration >= duration {
					top.Duration -= duration
					if top.Duration > 0.0 {
						(*bpmDurationStack)[len(*bpmDurationStack)-1] = top
					} else {
						*bpmDurationStack = (*bpmDurationStack)[:len(*bpmDurationStack)-1]
					}
					break
				} else {
					duration -= top.Duration
					*bpmDurationStack = (*bpmDurationStack)[:len(*bpmDurationStack)-1]
				}
			}
		}
	}
}

// Count 谱面数据统计
func (c *Chart) Count() *dto.ChartStats {
	var stats dto.ChartStats = dto.ChartStats{
		Time:    0.0,
		Notes:   0,
		BPMs:    make([]float64, 0),
		MainBPM: 0.0,
	}

	// 扁平化音符列表
	notes := c.flatten()

	// 计算 BPM 数
	var bpmCount int = 0
	for _, note := range notes {
		if note.Type == dto.NoteTypeBPM && note.BPM != 0.0 {
			bpmCount++
		}
	}
	var bpmDurationStack []bpmDuration = make([]bpmDuration, 0, bpmCount+1)

	var prevBeat float64 = 0.0
	var prevBPM float64 = 0.0
	var bpms map[float64]struct{} = make(map[float64]struct{}, bpmCount)
	// 再次遍历音符列表进行统计
	for _, note := range notes {
		if note.Type == dto.NoteTypeBPM {
			if prevBPM != 0.0 {
				// bpmDurationStack 为空与 prevBPM == 0.0 必定同时成立
				duration := (note.Beat - prevBeat) * 60.0 / math.Abs(prevBPM)
				handleBPMDuration(duration, prevBPM, &bpmDurationStack)
			}
			prevBPM = note.BPM
			prevBeat = note.Beat
			if prevBPM > 0.0 {
				if _, ok := bpms[prevBPM]; !ok {
					bpms[prevBPM] = struct{}{}
					stats.BPMs = append(stats.BPMs, prevBPM)
				}
			}
		} else {
			if !note.Hidden {
				stats.Notes++
			}
		}
	}
	// 统计收尾 BPM 时长
	if prevBPM != 0.0 && prevBeat != 0.0 {
		// 以谱面最后一个音符 Beat 值作为最后一个 BPM 计算的结束点
		duration := (notes[len(notes)-1].Beat - prevBeat) * 60.0 / math.Abs(prevBPM)
		handleBPMDuration(duration, prevBPM, &bpmDurationStack)
		if prevBPM > 0.0 && duration > 0.0 {
			if _, ok := bpms[prevBPM]; !ok {
				bpms[prevBPM] = struct{}{}
				stats.BPMs = append(stats.BPMs, prevBPM)
			}
		}
	}

	// 处理 BPM 计算栈计算谱面时长与 BPM 时长统计
	var bpmTimeMap map[float64]float64 = make(map[float64]float64, len(bpmDurationStack))
	var totalTime float64 = 0.0
	for _, bpmDur := range bpmDurationStack {
		totalTime += bpmDur.Duration
		if val, ok := bpmTimeMap[bpmDur.BPM]; ok {
			bpmTimeMap[bpmDur.BPM] = val + bpmDur.Duration
		} else {
			bpmTimeMap[bpmDur.BPM] = bpmDur.Duration
		}
	}
	stats.Time = totalTime

	// 计算主 BPM
	var mainDuration float64 = 0.0
	for bpm, duration := range bpmTimeMap {
		if duration > mainDuration {
			mainDuration = duration
			stats.MainBPM = bpm
		}
	}

	return &stats
}

// Standardize 对谱面进行标准化处理
func (c *Chart) Standardize() {
	sort.Sort(c)

	// 计算偏移量
	var offset float64 = 0.0
	for _, note := range *c {
		if note.Type == dto.NoteTypeBPM {
			offset = note.Beat
			break
		}
	}

	var result Chart = make(Chart, 0, len(*c))
	for _, note := range *c {
		if note.Type == dto.NoteTypeSlide || note.Type == dto.NoteTypeLong {
			var connections []Connection = make([]Connection, 0, len(note.Connections))

			// 统一类型名称
			note.Type = dto.NoteTypeSlide
			for index, connection := range note.Connections {
				// 修正偏移量
				connection.Beat -= offset
				if connection.Beat < 0.0 {
					connection.Beat = 0.0
				}
				// 修正节点字段值
				if index != 0 && index != len(note.Connections)-1 {
					connection.Flick = false
					connection.Skill = false
				}
				connections = append(connections, connection)
			}
			note.Connections = connections
		} else {
			// 修正偏移量
			note.Beat -= offset
			if note.Beat < 0.0 {
				note.Beat = 0.0
			}
		}
		result = append(result, note)
	}

	*c = result
}
