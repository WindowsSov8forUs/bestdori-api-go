package missions

import (
	"strconv"

	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/post"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// GetAll0 获取总任务 ID
func GetAll0(api *uniapi.UniAPI) (*dto.EmptyStruct, error) {
	endpoint := endpoints.MissionsAll(0)
	return uniapi.Get[dto.EmptyStruct](api, endpoint, nil)
}

// GetAll5 获取总任务详细信息
func GetAll5(api *uniapi.UniAPI) (*dto.EmptyStruct, error) {
	endpoint := endpoints.MissionsAll(5)
	return uniapi.Get[dto.EmptyStruct](api, endpoint, nil)
}

// Mission 任务
type Mission struct {
	Id   int
	Info *dto.MissionInfo
	api  *uniapi.UniAPI
}

// GetMission 获取任务实例
func GetMission(api *uniapi.UniAPI, id int) (*Mission, error) {
	endpoint := endpoints.MissionsInfo(id)
	if info, err := uniapi.Get[dto.MissionInfo](api, endpoint, nil); err != nil {
		return nil, err
	} else {
		return &Mission{
			Id:   id,
			Info: info,
			api:  api,
		}, nil
	}
}

func (m *Mission) Names() []*string {
	return m.Info.Title
}

func (m *Mission) DefaultServer() dto.ServerName {
	startAt := m.Info.StartAt
	switch {
	case startAt[0] != nil:
		return dto.ServerNameJP
	case startAt[1] != nil:
		return dto.ServerNameEN
	case startAt[2] != nil:
		return dto.ServerNameTW
	case startAt[3] != nil:
		return dto.ServerNameCN
	case startAt[4] != nil:
		return dto.ServerNameKR
	default:
		return ""
	}
}

// GetComments 获取任务评论
func (m *Mission) GetComments(limit, offset int, order post.Order) (*dto.PostList, error) {
	return post.GetList(
		m.api,
		"", false,
		"MISSION_COMMENT",
		strconv.Itoa(m.Id),
		nil, "",
		order,
		limit,
		offset,
	)
}
