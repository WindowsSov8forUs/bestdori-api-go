package miracleticket

import (
	"strconv"

	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// GetAll5 获取总自选券信息
func GetAll5(api *uniapi.UniAPI) (*dto.MiracleTicketExchangesAll5, error) {
	endpoint := endpoints.MiracleTicketExchangesAll(5)
	return uniapi.Get[dto.MiracleTicketExchangesAll5](api, endpoint, nil)
}

// MiracleTicketExchange 自选券
type MiracleTicketExchange struct {
	Id   int
	Info *dto.MiracleTicketExchangeInfo
	api  *uniapi.UniAPI
}

// GetMiracleTicketExchange 获取自选券实例
func GetMiracleTicketExchange(api *uniapi.UniAPI, id int) (*MiracleTicketExchange, error) {
	all, err := GetAll5(api)
	if err != nil {
		return nil, err
	}
	if info, ok := (*all)[strconv.Itoa(id)]; ok {
		return &MiracleTicketExchange{
			Id:   id,
			Info: &info,
			api:  api,
		}, nil
	} else {
		return nil, &bestdori.NotExistError{Target: "event archive " + strconv.Itoa(id)}
	}
}

func (mte *MiracleTicketExchange) Names() []*string {
	return mte.Info.Name
}

func (mte *MiracleTicketExchange) DefaultServer() dto.ServerName {
	ids := mte.Info.Ids
	switch {
	case ids[0] != nil:
		return dto.ServerNameJP
	case ids[1] != nil:
		return dto.ServerNameEN
	case ids[2] != nil:
		return dto.ServerNameTW
	case ids[3] != nil:
		return dto.ServerNameCN
	case ids[4] != nil:
		return dto.ServerNameKR
	default:
		return ""
	}
}

// GetIds 获取自选券包含的卡牌 ID 列表
func (mte *MiracleTicketExchange) GetIds(server dto.ServerName) ([]int, error) {
	ids := mte.Info.Ids[int(server.Id())]
	if ids == nil {
		return nil, &bestdori.NotExistError{
			Target: "miracle ticket exchange for server " + string(server),
		}
	}
	return *ids, nil
}
