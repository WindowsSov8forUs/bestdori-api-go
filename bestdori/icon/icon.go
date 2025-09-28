package icon

import (
	"strconv"

	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// GetBand 获取乐队图标
func GetBand(api *uniapi.UniAPI, id int) (*[]byte, error) {
	name := "band_" + strconv.Itoa(id)
	endpoint := endpoints.ResIconSvg(name)
	return uniapi.Get[[]byte](api, endpoint, nil)
}

// GetServer 获取服务器图标
func GetServer(api *uniapi.UniAPI, server dto.ServerName) (*[]byte, error) {
	endpoint := endpoints.ResIconPng(string(server))
	return uniapi.Get[[]byte](api, endpoint, nil)
}
