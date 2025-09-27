package eventtop

import (
	"fmt"

	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// GetData 获取活动 T10 排名分数线
//
// 当 interval 不为 0 时，获取最新数据
//
// 当 latest 为 1 时，获取最终数据
func GetData(
	api *uniapi.UniAPI,
	server dto.Server,
	event int,
	mid int,
	interval *int,
	latest *int,
) (*dto.EventTopData, error) {
	var params = map[string]interface{}{
		"server": server,
		"event":  event,
		"mid":    mid,
	}
	if interval != nil {
		params["interval"] = *interval
	}
	if latest != nil {
		params["latest"] = *latest
	}
	result, err := uniapi.Get[dto.EventTopData](api, endpoints.TrackerEventTop, params)
	if err == nil {
		return result, nil
	} else if err.(*uniapi.ResponseStatusError).StatusCode() == 404 {
		return nil, &bestdori.NotExistError{Target: fmt.Sprintf("event %d", event)}
	}
	return nil, err
}
