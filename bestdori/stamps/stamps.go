package stamps

import (
	"strconv"

	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// GetAll 获取总贴纸信息
func GetAll2(api *uniapi.UniAPI) (*dto.StampsAll2, error) {
	endpoint := endpoints.StampsAll(2)
	return uniapi.Get[dto.StampsAll2](api, endpoint, nil)
}

// Stamp 贴纸
type Stamp struct {
	Id   int
	Info *dto.StampInfo
	api  *uniapi.UniAPI
}

// GetStamp 获取贴纸实例
func GetStamp(api *uniapi.UniAPI, stampId int) (*Stamp, error) {
	all, err := GetAll2(api)
	if err != nil {
		return nil, err
	}
	if info, ok := (*all)[strconv.Itoa(stampId)]; ok {
		return &Stamp{
			Id:   stampId,
			Info: &info,
			api:  api,
		}, nil
	}
	return nil, &bestdori.NotExistError{Target: "stamp " + strconv.Itoa(stampId)}
}

// GetImage 获取贴纸图片
func (s *Stamp) GetImage(server dto.ServerName) (*[]byte, error) {
	endpoint := endpoints.StampGet(string(server), s.Info.ImageName)
	return uniapi.Get[[]byte](s.api, endpoint, nil)
}
