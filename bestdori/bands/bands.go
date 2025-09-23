package bands

import (
	"fmt"

	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// GetAll 获取所有乐队信息
func GetAll(api *uniapi.UniAPI) (*dto.BandsAll1, error) {
	endpoint := fmt.Sprintf(endpoints.BandsAll, 1)
	return uniapi.Get[dto.BandsAll1](api, endpoint, nil)
}

// GetMain 获取主要乐队信息
func GetMain(api *uniapi.UniAPI) (*dto.BandsMain1, error) {
	endpoint := fmt.Sprintf(endpoints.BandsMain, 1)
	return uniapi.Get[dto.BandsMain1](api, endpoint, nil)
}

// GetLogo 获取乐队 Logo 图片数据
func GetLogo(api *uniapi.UniAPI, id int, typ, server string) (*[]byte, error) {
	endpoint := fmt.Sprintf(endpoints.BandLogo, server, id, typ)
	return uniapi.Get[[]byte](api, endpoint, nil)
}
