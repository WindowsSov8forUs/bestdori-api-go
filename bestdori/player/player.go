package player

import (
	"strconv"

	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// Player 玩家
type Player struct {
	Id      int
	Server  dto.ServerName
	Profile *dto.PlayerDataProfile
	api     *uniapi.UniAPI
}

// GetPlayer 获取玩家实例
func GetPlayer(api *uniapi.UniAPI, id int, server dto.ServerName, mode uint8) (*Player, error) {
	params := map[string]any{
		"mode": mode,
	}
	endpoints := endpoints.PlayerInfo(string(server), id)
	result, err := uniapi.Get[dto.PlayerInfo](api, endpoints, params)
	if err != nil {
		return nil, err
	}

	if profile := result.Data.Profile; profile != nil {
		return &Player{
			Id:      id,
			Server:  server,
			Profile: profile,
			api:     api,
		}, nil
	} else {
		return nil, &bestdori.NotExistError{
			Target: "player " + string(server) + " " + strconv.Itoa(id),
		}
	}
}
