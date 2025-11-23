package events

import (
	"fmt"
	"strconv"

	"github.com/samber/lo"

	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	eventarchive "github.com/WindowsSov8forUs/bestdori-api-go/bestdori/event_archive"
	eventtracker "github.com/WindowsSov8forUs/bestdori-api-go/bestdori/event_tracker"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/festival"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/post"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/stamps"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// GetAll0 获取总活动 ID
func GetAll0(api *uniapi.UniAPI) (*dto.EmptyStruct, error) {
	endpoint := endpoints.EventsAll(0)
	return uniapi.Get[dto.EmptyStruct](api, endpoint, nil)
}

// GetAll1 获取总活动 ID 和名称
func GetAll1(api *uniapi.UniAPI) (*dto.EventsAll1, error) {
	endpoint := endpoints.EventsAll(1)
	return uniapi.Get[dto.EventsAll1](api, endpoint, nil)
}

// GetAll3 获取总活动简介信息
func GetAll3(api *uniapi.UniAPI) (*dto.EventsAll3, error) {
	endpoint := endpoints.EventsAll(3)
	return uniapi.Get[dto.EventsAll3](api, endpoint, nil)
}

// GetAll4 获取总活动较详细信息
func GetAll4(api *uniapi.UniAPI) (*dto.EventsAll4, error) {
	endpoint := endpoints.EventsAll(4)
	return uniapi.Get[dto.EventsAll4](api, endpoint, nil)
}

// GetAll5 获取总活动详细信息
func GetAll5(api *uniapi.UniAPI) (*dto.EventArchiveAll5, error) {
	endpoint := endpoints.EventsAll(5)
	return uniapi.Get[dto.EventArchiveAll5](api, endpoint, nil)
}

// GetAll6 获取总活动详细信息
func GetAll6(api *uniapi.UniAPI) (*dto.EventsAll6, error) {
	endpoint := endpoints.EventsAll(6)
	return uniapi.Get[dto.EventsAll6](api, endpoint, nil)
}

// Event 活动
type Event struct {
	Id      int
	Info    *dto.EventInfo
	Archive *eventarchive.EventArchive
	api     *uniapi.UniAPI
}

// GetEvent 获取活动实例
func GetEvent(api *uniapi.UniAPI, eventId int) (*Event, error) {
	endpoint := endpoints.EventsInfo(eventId)
	info, err := uniapi.Get[dto.EventInfo](api, endpoint, nil)
	if err != nil {
		return nil, err
	}

	archive, err := eventarchive.GetEventArchive(api, eventId)
	if err != nil {
		return nil, err
	}

	return &Event{
		Id:      eventId,
		Info:    info,
		Archive: archive,
		api:     api,
	}, nil
}

// GetComments 获取活动评论
func (e *Event) GetComments(limit, offset int, order post.Order) (*dto.PostList, error) {
	return post.GetList(
		e.api,
		"", false,
		"EVENT_COMMENT",
		strconv.Itoa(e.Id),
		nil, "",
		order,
		limit,
		offset,
	)
}

// Tracker 获取对应服务器活动PT&排名追踪器
func (e *Event) Tracker(server dto.Server) *eventtracker.EventTracker {
	return eventtracker.GetEventTracker(e.api, server, e.Id)
}

// GetBanner 获取活动缩略图图像
func (e *Event) GetBanner(server dto.ServerName) (*[]byte, error) {
	// 判断服务器
	startAt := e.Info.StartAt
	serverId := server.Id()
	if startAt[serverId] == nil {
		return nil, &bestdori.ServerNotAvailableError{
			Target: "event " + strconv.Itoa(e.Id),
			Server: server,
		}
	}

	endpoint := endpoints.EventBanner(string(server), e.Info.AssetBundleName)
	return uniapi.Get[[]byte](e.api, endpoint, nil)
}

// GetLogo 获取活动 Logo 图像
func (e *Event) GetLogo(server dto.ServerName) (*[]byte, error) {
	// 判断服务器
	startAt := e.Info.StartAt
	serverId := server.Id()
	if startAt[serverId] == nil {
		return nil, &bestdori.ServerNotAvailableError{
			Target: "event " + strconv.Itoa(e.Id),
			Server: server,
		}
	}

	endpoint := endpoints.EventLogo(string(server), e.Info.AssetBundleName)
	return uniapi.Get[[]byte](e.api, endpoint, nil)
}

// GetTopScreen 获取活动主界面图像
func (e *Event) GetTopScreen(server dto.ServerName, typ string) (*[]byte, error) {
	// 判断服务器
	startAt := e.Info.StartAt
	serverId := server.Id()
	if startAt[serverId] == nil {
		return nil, &bestdori.ServerNotAvailableError{
			Target: "event " + strconv.Itoa(e.Id),
			Server: server,
		}
	}

	endpoint := endpoints.EventTopScreen(string(server), e.Info.AssetBundleName, typ)
	return uniapi.Get[[]byte](e.api, endpoint, nil)
}

// GetStamp 获取活动奖励稀有表情
func (e *Event) GetStamp() (*[]byte, error) {
	// 获取第一个非 nil 奖励列表
	pointReward, ok := lo.Find(e.Info.PointRewards, func(pr *[]dto.EventPointReward) bool {
		return pr != nil
	})
	if !ok {
		return nil, &bestdori.NotExistError{Target: "point reward of event " + strconv.Itoa(e.Id)}
	}

	// 获取 rewardType 为 stamp 的活动奖励
	reward, ok := lo.Find(*pointReward, func(r dto.EventPointReward) bool {
		return r.RewardType == "stamp"
	})
	if !ok {
		return nil, &bestdori.NotExistError{Target: "stamp reward of event " + strconv.Itoa(e.Id)}
	}
	if reward.RewardId == nil {
		return nil, &bestdori.NotExistError{Target: "stamp reward of event " + strconv.Itoa(e.Id)}
	}

	stamp, err := stamps.GetStamp(e.api, *reward.RewardId)
	if err != nil {
		return nil, err
	}
	return stamp.GetImage(dto.ServerNameJP)
}

// GetTop 获取 T10 排名分数线
func (e *Event) GetTop(server dto.Server, mid int, interval *int, latest *int) (*dto.EventTopData, error) {
	if interval == nil && latest == nil {
		return nil, fmt.Errorf("either `interval` or `latest` must be provided")
	}
	if interval == nil {
		return e.Archive.GetTop(server, mid)
	} else {
		return e.Tracker(server).GetTop(mid, interval)
	}
}

// GetRotationMusics 获取团队 LIVE 佳节活动歌曲循环数据
func (e *Event) GetRotationMusics() (*[]dto.FestivalRotationMusic, error) {
	if e.Info.EventType != "festival" {
		return nil, fmt.Errorf("rotation musics are only available for festival events, not `%s`", e.Info.EventType)
	}
	return festival.GetRotationMusics(e.api, e.Id)
}

// GetStages 获取团队 LIVE 佳节活动舞台数据
func (e *Event) GetStages() (*[]dto.FestivalStage, error) {
	if e.Info.EventType != "festival" {
		return nil, fmt.Errorf("stages are only available for festival events, not `%s`", e.Info.EventType)
	}
	return festival.GetStages(e.api, e.Id)
}
