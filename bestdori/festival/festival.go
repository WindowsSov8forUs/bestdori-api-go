package festival

import (
	"fmt"

	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// GetRotationMusics 获取歌曲循环数据
func GetRotationMusics(api *uniapi.UniAPI, id int) (*[]dto.FestivalRotationMusic, error) {
	endpoint := fmt.Sprintf(endpoints.FestivalRotationMusics, id)
	return uniapi.Get[[]dto.FestivalRotationMusic](api, endpoint, nil)
}

// GetStages 获取活动舞台数据
func GetStages(api *uniapi.UniAPI, id int) (*[]dto.FestivalStage, error) {
	endpoint := fmt.Sprintf(endpoints.FestivalStages, id)
	return uniapi.Get[[]dto.FestivalStage](api, endpoint, nil)
}
