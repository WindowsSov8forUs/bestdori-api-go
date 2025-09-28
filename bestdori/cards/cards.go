package cards

import (
	"fmt"
	"strconv"

	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/post"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// GetAll0 获取总卡牌 ID
func GetAll0(api *uniapi.UniAPI) (*dto.EmptyStruct, error) {
	endpoint := endpoints.CardsAll(0)
	return uniapi.Get[dto.EmptyStruct](api, endpoint, nil)
}

// GetAll2 获取总卡牌属性信息
func GetAll2(api *uniapi.UniAPI) (*dto.CardsAll2, error) {
	endpoint := endpoints.CardsAll(2)
	return uniapi.Get[dto.CardsAll2](api, endpoint, nil)
}

// GetAll3 获取总卡牌简洁信息
func GetAll3(api *uniapi.UniAPI) (*dto.CardsAll3, error) {
	endpoint := endpoints.CardsAll(3)
	return uniapi.Get[dto.CardsAll3](api, endpoint, nil)
}

// GetAll5 获取总卡牌较详细信息
func GetAll5(api *uniapi.UniAPI) (*dto.CardsAll5, error) {
	endpoint := endpoints.CardsAll(5)
	return uniapi.Get[dto.CardsAll5](api, endpoint, nil)
}

// GetAttributeIcon 获取卡牌属性图标
func GetAttributeIcon(api *uniapi.UniAPI, attribute dto.CardAttribute) (*[]byte, error) {
	endpoint := endpoints.ResIconSvg(string(attribute))
	return uniapi.Get[[]byte](api, endpoint, nil)
}

// GetStarIcon 获取星星图标
func GetStarIcon(api *uniapi.UniAPI, star dto.StarType) (*[]byte, error) {
	endpoint := endpoints.ResIconPng(string(star))
	return uniapi.Get[[]byte](api, endpoint, nil)
}

// GetFrame 获取卡牌完整边框
func GetFrame(api *uniapi.UniAPI, rarity dto.CardRarity, attribute *dto.CardAttribute) (*[]byte, error) {
	var name string
	if rarity == dto.CardRarity1 {
		if attribute == nil {
			return nil, fmt.Errorf("param `attribute` is required when rarity is 1")
		}
		name = "frame-" + strconv.Itoa(int(rarity)) + "-" + string(*attribute)
	} else {
		name = "frame-" + strconv.Itoa(int(rarity))
	}
	endpoint := endpoints.ResImagePng(name)
	return uniapi.Get[[]byte](api, endpoint, nil)
}

// GetCardFrame 获取卡牌缩略图边框
func GetCardFrame(api *uniapi.UniAPI, rarity dto.CardRarity, attribute *dto.CardAttribute) (*[]byte, error) {
	var name string
	if rarity == dto.CardRarity1 {
		if attribute == nil {
			return nil, fmt.Errorf("param `attribute` is required when rarity is 1")
		}
		name = "card-" + strconv.Itoa(int(rarity)) + "-" + string(*attribute)
	} else {
		name = "card-" + strconv.Itoa(int(rarity))
	}
	endpoint := endpoints.ResImagePng(name)
	return uniapi.Get[[]byte](api, endpoint, nil)
}

// Card 卡牌
type Card struct {
	Id   int
	Info *dto.CardInfo
	api  *uniapi.UniAPI
}

// GetCard 获取卡牌实例
func GetCard(api *uniapi.UniAPI, id int) (*Card, error) {
	endpoint := endpoints.CardsInfo(id)
	info, err := uniapi.Get[dto.CardInfo](api, endpoint, nil)
	if err != nil {
		return nil, err
	}
	return &Card{
		Id:   id,
		Info: info,
		api:  api,
	}, nil
}

func (c *Card) Names() []*string {
	return c.Info.Prefix
}

func (c *Card) DefaultServer() dto.ServerName {
	releasedAt := c.Info.ReleasedAt
	switch {
	case releasedAt[0] != nil:
		return dto.ServerNameJP
	case releasedAt[1] != nil:
		return dto.ServerNameEN
	case releasedAt[2] != nil:
		return dto.ServerNameTW
	case releasedAt[3] != nil:
		return dto.ServerNameCN
	case releasedAt[4] != nil:
		return dto.ServerNameKR
	default:
		return ""
	}
}

// GetComments 获取卡牌评论
func (c *Card) GetComments(limit, offset int, order post.Order) (*dto.PostList, error) {
	categoryName := "CARD_COMMENT"
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

// GetImage 获取卡牌完整图片
func (c *Card) GetImage(typ dto.CardTrain) (*[]byte, error) {
	endpoint := endpoints.CharactersResourceSet(
		string(c.DefaultServer()), c.Info.ResourceSetName, "card", string(typ),
	)
	return uniapi.Get[[]byte](c.api, endpoint, nil)
}

// GetTrim 获取卡牌无背景图片
func (c *Card) GetTrim(typ dto.CardTrain) (*[]byte, error) {
	endpoint := endpoints.CharactersResourceSet(
		string(c.DefaultServer()), c.Info.ResourceSetName, "trim", string(typ),
	)
	return uniapi.Get[[]byte](c.api, endpoint, nil)
}

// GetThumb 获取卡牌缩略图图片
func (c *Card) GetThumb(typ dto.CardTrain) (*[]byte, error) {
	endpoint := endpoints.ThumbChara(
		string(c.DefaultServer()), c.Id/50, c.Info.ResourceSetName, string(typ),
	)
	return uniapi.Get[[]byte](c.api, endpoint, nil)
}

// GetLiveSD 获取 LIVE 服装图片
func (c *Card) GetLiveSD() (*[]byte, error) {
	endpoint := endpoints.CharactersLiveSD(
		string(c.DefaultServer()), c.Info.ResourceSetName,
	)
	return uniapi.Get[[]byte](c.api, endpoint, nil)
}
