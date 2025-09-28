package costumes

import (
	"strconv"

	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/post"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// GetAll0 获取总服装 ID
func GetAll0(api *uniapi.UniAPI) (*dto.EmptyStruct, error) {
	endpoint := endpoints.CostumesAll(0)
	return uniapi.Get[dto.EmptyStruct](api, endpoint, nil)
}

// GetAll5 获取总服装信息
func GetAll5(api *uniapi.UniAPI) (*dto.CostumesAll5, error) {
	endpoint := endpoints.CostumesAll(5)
	return uniapi.Get[dto.CostumesAll5](api, endpoint, nil)
}

// Costume 服装
type Costume struct {
	Id   int
	Info *dto.CostumeInfo
	api  *uniapi.UniAPI
}

// GetCostume 获取服装实例
func GetCostume(api *uniapi.UniAPI, id int) (*Costume, error) {
	endpoint := endpoints.CostumesInfo(id)
	info, err := uniapi.Get[dto.CostumeInfo](api, endpoint, nil)
	if err != nil {
		return nil, err
	}
	return &Costume{
		Id:   id,
		Info: info,
		api:  api,
	}, nil
}

func (c *Costume) Names() []*string {
	return c.Info.Description
}

func (c *Costume) DefaultServer() dto.ServerName {
	publishedAt := c.Info.PublishedAt
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

// GetComments 获取服装评论
func (c *Costume) GetComments(limit, offset int, order post.Order) (*dto.PostList, error) {
	categoryName := "COSTUME_COMMENT"
	categoryId := strconv.Itoa(c.Id)

	return post.GetList(
		c.api,
		nil, nil,
		&categoryName,
		&categoryId,
		nil, nil,
		order,
		limit,
		offset,
	)
}

// GetSdchara 获取 LIVE 服装图片
func (c *Costume) GetSdchara() (*[]byte, error) {
	endpoint := endpoints.CharactersLiveSD(string(c.DefaultServer()), c.Info.SdResourceName)
	return uniapi.Get[[]byte](c.api, endpoint, nil)
}

// GetIcon 获取图标
func (c *Costume) GetIcon() (*[]byte, error) {
	endpoint := endpoints.ThumbCostume(string(c.DefaultServer()), c.Id/50, c.Info.AssetBundleName)
	return uniapi.Get[[]byte](c.api, endpoint, nil)
}
