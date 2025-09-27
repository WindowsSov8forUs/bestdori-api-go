package costumes

import (
	"fmt"

	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/post"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// GetAll0 获取总服装 ID
func GetAll0(api *uniapi.UniAPI) (*dto.EmptyStruct, error) {
	endpoint := fmt.Sprintf(endpoints.CostumesAll, 0)
	return uniapi.Get[dto.EmptyStruct](api, endpoint, nil)
}

// GetAll5 获取总服装信息
func GetAll5(api *uniapi.UniAPI) (*dto.CostumesAll5, error) {
	endpoint := fmt.Sprintf(endpoints.CostumesAll, 5)
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
	endpoint := fmt.Sprintf(endpoints.CostumesInfo, id)
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
	if publishedAt[0] != nil {
		return dto.ServerNameJP
	} else if publishedAt[1] != nil {
		return dto.ServerNameEN
	} else if publishedAt[2] != nil {
		return dto.ServerNameTW
	} else if publishedAt[3] != nil {
		return dto.ServerNameCN
	} else if publishedAt[4] != nil {
		return dto.ServerNameKR
	} else {
		return ""
	}
}

// GetComments 获取服装评论
func (c *Costume) GetComments(limit, offset int, order post.Order) (*dto.PostList, error) {
	categoryName := "COSTUME_COMMENT"
	categoryId := fmt.Sprintf("%d", c.Id)

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
	endpoint := fmt.Sprintf(endpoints.CharactersLiveSD, c.DefaultServer(), c.Info.SdResourceName)
	return uniapi.Get[[]byte](c.api, endpoint, nil)
}

// GetIcon 获取图标
func (c *Costume) GetIcon() (*[]byte, error) {
	endpoint := fmt.Sprintf(
		endpoints.ThumbCostume,
		c.DefaultServer(), c.Id/50, c.Info.AssetBundleName,
	)
	return uniapi.Get[[]byte](c.api, endpoint, nil)
}
