package eventarchive

import (
	"strconv"

	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	eventtop "github.com/WindowsSov8forUs/bestdori-api-go/bestdori/event_top"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/post"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// GetAll5 获取总活动数据信息
func GetAll5(api *uniapi.UniAPI) (*dto.EventArchiveAll5, error) {
	endpoint := endpoints.ArchivesAll(5)
	return uniapi.Get[dto.EventArchiveAll5](api, endpoint, nil)
}

// EventArchive 活动数据
type EventArchive struct {
	Id   int
	Info *dto.EventArchiveInfo
	api  *uniapi.UniAPI
}

// GetEventArchive 获取活动数据实例
func GetEventArchive(api *uniapi.UniAPI, id int) (*EventArchive, error) {
	all, err := GetAll5(api)
	if err != nil {
		return nil, err
	}
	if info, ok := (*all)[strconv.Itoa(id)]; ok {
		return &EventArchive{
			Id:   id,
			Info: &info,
			api:  api,
		}, nil
	} else {
		return nil, &bestdori.NotExistError{Target: "event archive " + strconv.Itoa(id)}
	}
}

// GetComments 获取活动数据评论
func (ea *EventArchive) GetComments(limit, offset int, order post.Order) (*dto.PostList, error) {
	return post.GetList(
		ea.api,
		"", false,
		"EVENTARCHIVE_COMMENT",
		strconv.Itoa(ea.Id),
		nil, "",
		order,
		limit,
		offset,
	)
}

// GetTop 获取最终排名分数线
func (ea *EventArchive) GetTop(server dto.Server, mid int) (*dto.EventTopData, error) {
	var latest int = 1
	return eventtop.GetData(ea.api, server, ea.Id, mid, nil, &latest)
}
