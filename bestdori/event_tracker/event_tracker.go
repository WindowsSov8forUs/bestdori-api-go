package eventtracker

import (
	"fmt"

	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	eventtop "github.com/WindowsSov8forUs/bestdori-api-go/bestdori/event_top"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/post"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// GetRates 获取活动追踪比率数组
func GetRates(api *uniapi.UniAPI) (*[]dto.EventTrackerRate, error) {
	return uniapi.Get[[]dto.EventTrackerRate](api, endpoints.TrackerRates, nil)
}

// EventTracker 活动排名追踪器
type EventTracker struct {
	Server dto.Server
	Event  int
	api    *uniapi.UniAPI
}

// GetEventTracker 获取活动排名追踪器
func GetEventTracker(api *uniapi.UniAPI, server dto.Server, event int) *EventTracker {
	return &EventTracker{
		Server: server,
		Event:  event,
		api:    api,
	}
}

// GetComments 获取活动排名追踪评论
func (et *EventTracker) GetComments(limit, offset int, order post.Order) (*dto.PostList, error) {
	categoryName := "EVENTTRACKER_COMMENT"
	categoryId := fmt.Sprintf("%d", et.Event)

	return post.GetList(
		et.api,
		nil, nil,
		&categoryName,
		&categoryId,
		nil, nil,
		order,
		limit,
		offset,
	)
}

// GetTop 获取 T10 实时排名追踪信息
func (et *EventTracker) GetTop(mid int, interval *int) (*dto.EventTopData, error) {
	return eventtop.GetData(
		et.api, et.Server, et.Event, mid, interval, nil,
	)
}

// GetData 获取分数线追踪信息
func (et *EventTracker) GetData(tier int) (*dto.EventTrackerData, error) {
	var params = map[string]interface{}{
		"server": et.Server,
		"event":  et.Event,
		"tier":   tier,
	}
	result, err := uniapi.Get[dto.EventTrackerData](et.api, endpoints.TrackerEventTracker, params)
	if err == nil {
		return result, nil
	} else if err.(*uniapi.ResponseStatusError).StatusCode() == 404 {
		return nil, &bestdori.NotExistError{Target: fmt.Sprintf("event %d", et.Event)}
	}
	return nil, err
}
