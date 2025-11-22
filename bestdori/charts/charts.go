package charts

import (
	"github.com/mitchellh/mapstructure"

	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// GetChart 获取官方谱面
func GetChart(api *uniapi.UniAPI, id int, difficulty dto.ChartDifficultyName) (*dto.Chart, error) {
	endpoint := endpoints.ChartsInfo(id, string(difficulty))
	return uniapi.Get[dto.Chart](api, endpoint, nil)
}

func unmarshalMap(data map[string]any) (*dto.Note, error) {
	var note dto.Note
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

// UnmarshalSlice 从 []map[string]any 中解析谱面
func UnmarshalSlice(data []map[string]any) (*dto.Chart, error) {
	var chart dto.Chart = make(dto.Chart, 0, len(data))
	for _, item := range data {
		note, err := unmarshalMap(item)
		if err != nil {
			return nil, err
		}
		chart = append(chart, *note)
	}
	return &chart, nil
}
