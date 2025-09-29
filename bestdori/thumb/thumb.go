package thumb

import (
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// GetChara 获取卡牌缩略图
func GetChara(api *uniapi.UniAPI, id int, resourceSetName string, typ dto.CardTrain, server dto.ServerName) (*[]byte, error) {
	endpoint := endpoints.ThumbChara(
		string(server), id/50, resourceSetName, string(typ),
	)
	return uniapi.Get[[]byte](api, endpoint, nil)
}

// GetDegree 获取称号资源
func GetDegree(api *uniapi.UniAPI, degreeName string, server dto.ServerName) (*[]byte, error) {
	endpoint := endpoints.ThumbDegree(string(server), degreeName)
	return uniapi.Get[[]byte](api, endpoint, nil)
}

// GetCostume 获取服装图标
func GetCostume(api *uniapi.UniAPI, id int, assetBundleName string, server dto.ServerName) (*[]byte, error) {
	endpoint := endpoints.ThumbCostume(string(server), id/50, assetBundleName)
	return uniapi.Get[[]byte](api, endpoint, nil)
}
