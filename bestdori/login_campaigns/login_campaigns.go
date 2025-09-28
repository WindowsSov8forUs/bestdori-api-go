package logincampaigns

import (
	"strconv"

	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/post"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// GetAll0 获取总登录奖励 ID
func GetAll0(api *uniapi.UniAPI) (*dto.EmptyStruct, error) {
	endpoint := endpoints.LoginCampaignsAll(0)
	return uniapi.Get[dto.EmptyStruct](api, endpoint, nil)
}

// GetAll1 获取总登录奖励简洁信息
func GetAll1(api *uniapi.UniAPI) (*dto.LoginCampaignsAll1, error) {
	endpoint := endpoints.LoginCampaignsAll(1)
	return uniapi.Get[dto.LoginCampaignsAll1](api, endpoint, nil)
}

// GetAll5 获取总登录奖励详细信息
func GetAll5(api *uniapi.UniAPI) (*dto.LoginCampaignsAll5, error) {
	endpoint := endpoints.LoginCampaignsAll(5)
	return uniapi.Get[dto.LoginCampaignsAll5](api, endpoint, nil)
}

// LoginCampaign 登录奖励
type LoginCampaign struct {
	Id   int
	Info *dto.LoginCampaignInfo
	api  *uniapi.UniAPI
}

// GetLoginCampaign 获取登录奖励实例
func GetLoginCampaign(api *uniapi.UniAPI, id int) (*LoginCampaign, error) {
	endpoint := endpoints.LoginCampaignsInfo(id)
	if info, err := uniapi.Get[dto.LoginCampaignInfo](api, endpoint, nil); err != nil {
		return nil, err
	} else {
		return &LoginCampaign{
			Id:   id,
			Info: info,
			api:  api,
		}, nil
	}
}

func (lc *LoginCampaign) Names() []*string {
	return lc.Info.Caption
}

func (lc *LoginCampaign) DefaultServer() dto.ServerName {
	publishedAt := lc.Info.PublishedAt
	switch {
	case publishedAt[0] != nil:
		return dto.ServerNameJP
	case publishedAt[1] != nil:
		return dto.ServerNameEN
	case publishedAt[2] != nil:
		return dto.ServerNameTW
	case publishedAt[3] != nil:
		return dto.ServerNameCN
	case publishedAt[4] != nil:
		return dto.ServerNameKR
	default:
		return ""
	}
}

// GetComments 获取登录奖励评论
func (lc *LoginCampaign) GetComments(limit, offset int, order post.Order) (*dto.PostList, error) {
	categoryName := "LOGINCAMPAIGN_COMMENT"
	categoryId := strconv.Itoa(lc.Id)

	return post.GetList(
		lc.api,
		nil, nil,
		&categoryName,
		&categoryId,
		nil, nil,
		order,
		limit,
		offset,
	)
}

// GetBackground 获取登录奖励背景图
func (lc *LoginCampaign) GetBackground(server dto.ServerName) (*[]byte, error) {
	assetBundleName := lc.Info.AssetBundleName
	if assetBundleName[int(server.Id())] == nil {
		return nil, &bestdori.NotExistError{
			Target: "login campaign " + strconv.Itoa(lc.Id) + " in server " + string(server),
		}
	}

	endpoint := endpoints.EventLoginBonus(
		string(server), *assetBundleName[int(server.Id())],
	)
	return uniapi.Get[[]byte](lc.api, endpoint, nil)
}
