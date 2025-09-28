package gacha

import (
	"strconv"

	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/post"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// GetAll0 获取总招募 ID
func GetAll0(api *uniapi.UniAPI) (*dto.EmptyStruct, error) {
	endpoint := endpoints.GachaAll(0)
	return uniapi.Get[dto.EmptyStruct](api, endpoint, nil)
}

// GetAll1 获取总招募简洁信息
func GetAll1(api *uniapi.UniAPI) (*dto.GachaAll1, error) {
	endpoint := endpoints.GachaAll(1)
	return uniapi.Get[dto.GachaAll1](api, endpoint, nil)
}

// GetAll3 获取总招募较详细信息
func GetAll3(api *uniapi.UniAPI) (*dto.GachaAll3, error) {
	endpoint := endpoints.GachaAll(3)
	return uniapi.Get[dto.GachaAll3](api, endpoint, nil)
}

// GetAll5 获取总招募详细信息
func GetAll5(api *uniapi.UniAPI) (*dto.GachaAll5, error) {
	endpoint := endpoints.GachaAll(5)
	return uniapi.Get[dto.GachaAll5](api, endpoint, nil)
}

// Gacha 招募
type Gacha struct {
	Id   int
	Info *dto.GachaInfo
	api  *uniapi.UniAPI
}

// GetGacha 获取招募实例
func GetGacha(api *uniapi.UniAPI, id int) (*Gacha, error) {
	endpoint := endpoints.GachaInfo(id)
	if info, err := uniapi.Get[dto.GachaInfo](api, endpoint, nil); err != nil {
		return nil, err
	} else {
		return &Gacha{
			Id:   id,
			Info: info,
			api:  api,
		}, nil
	}
}

func (g *Gacha) Names() []*string {
	return g.Info.GachaName
}

func (g *Gacha) DefaultServer() dto.ServerName {
	publishedAt := g.Info.PublishedAt
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

// GetComments 获取招募评论
func (g *Gacha) GetComments(limit, offset int, order post.Order) (*dto.PostList, error) {
	categoryName := "GACHA_COMMENT"
	categoryId := strconv.Itoa(g.Id)

	return post.GetList(
		g.api,
		nil, nil,
		&categoryName,
		&categoryId,
		nil, nil,
		order,
		limit,
		offset,
	)
}

// GetBanner 获取招募缩略图图像
func (g *Gacha) GetBanner(server dto.ServerName) (*[]byte, error) {
	// 判断服务器
	publishedAt := g.Info.PublishedAt
	serverId := server.Id()
	if publishedAt[serverId] == nil {
		return nil, &bestdori.ServerNotAvailableError{
			Target: "gacha " + strconv.Itoa(g.Id),
			Server: server,
		}
	}

	endpoint := endpoints.HomebannerGet(string(server), g.Info.BannerAssetBundleName)
	return uniapi.Get[[]byte](g.api, endpoint, nil)
}

// GetPickups 获取招募 Pickup 图像
func (g *Gacha) GetPickups(server dto.ServerName) ([]*[]byte, error) {
	// 判断服务器
	publishedAt := g.Info.PublishedAt
	serverId := server.Id()
	if publishedAt[serverId] == nil {
		return nil, &bestdori.ServerNotAvailableError{
			Target: "gacha " + strconv.Itoa(g.Id),
			Server: server,
		}
	}

	var results []*[]byte = make([]*[]byte, 0, 3)
	var pickups = []string{"pickup1", "pickup2", "pickup"}
	for _, pickup := range pickups {
		endpoint := endpoints.GachaScreen(string(server), g.Id, pickup)
		if data, err := uniapi.Get[[]byte](g.api, endpoint, nil); err != nil {
			continue
		} else {
			results = append(results, data)
		}
	}
	if len(results) < 1 {
		return nil, &bestdori.NotExistError{Target: "gacha " + strconv.Itoa(g.Id) + " pickup image"}
	}
	return results, nil
}

// GetLogo 获取招募 Logo 图像
func (g *Gacha) GetLogo(server dto.ServerName) (*[]byte, error) {
	// 判断服务器
	publishedAt := g.Info.PublishedAt
	serverId := server.Id()
	if publishedAt[serverId] == nil {
		return nil, &bestdori.ServerNotAvailableError{
			Target: "gacha " + strconv.Itoa(g.Id),
			Server: server,
		}
	}

	endpoint := endpoints.GachaScreen(string(server), g.Id, "logo")
	return uniapi.Get[[]byte](g.api, endpoint, nil)
}
