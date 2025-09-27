package charts

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"

	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
)

type Connection struct {
	Beat   float64 `json:"beat"`             // 音符所在节拍值
	Lane   int     `json:"lane"`             // 音符所在轨道
	Flick  bool    `json:"flick,omitempty"`  // 是否为 flick 音符
	Skill  bool    `json:"skill,omitempty"`  // 是否为技能音符
	Hidden bool    `json:"hidden,omitempty"` // 是否为隐藏音符
}

type Note struct {
	Type        dto.NoteType      `json:"type"`                  // 音符类型
	Beat        float64           `json:"beat"`                  // 音符所在节拍值
	BPM         float64           `json:"bpm,omitempty"`         // 将要改变为的 BPM 值
	Connections []Connection      `json:"connections,omitempty"` // 连线音符列表
	Direction   dto.NoteDirection `json:"direction,omitempty"`   // 音符方向
	Flick       bool              `json:"flick,omitempty"`       // 是否为 flick 音符
	Skill       bool              `json:"skill,omitempty"`       // 是否为技能音符
	Hidden      bool              `json:"hidden,omitempty"`      // 是否为隐藏音符
	Lane        int               `json:"lane"`                  // 音符所在轨道
}

func unmarshalMap(data map[string]any) (*Note, error) {
	var note Note
	config := &mapstructure.DecoderConfig{
		TagName:    "json",
		Result:     &note,
		ZeroFields: true,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}
	if err := decoder.Decode(data); err != nil {
		return nil, err
	}
	return &note, nil
}

func (n *Note) MarshalJSON() ([]byte, error) {
	if n.Type != dto.NoteTypeBPM && n.Type != dto.NoteTypeSlide && n.Type != dto.NoteTypeLong {
		return json.Marshal(n)
	}
	type Alias Note
	if n.Type == dto.NoteTypeBPM {
		type Aux struct {
			Alias
			Lane int `json:"lane,omitempty"`
		}
		aux := Aux{Alias: Alias(*n), Lane: n.Lane}
		return json.Marshal(aux)
	} else {
		type Aux struct {
			Alias
			Beat float64 `json:"beat,omitempty"`
			Lane int     `json:"lane,omitempty"`
		}
		aux := Aux{Alias: Alias(*n), Beat: n.Beat, Lane: n.Lane}
		return json.Marshal(aux)
	}
}
