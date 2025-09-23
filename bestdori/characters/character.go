package characters

import (
	"fmt"

	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// GetAll0 获取总角色 ID
func GetAll0(api *uniapi.UniAPI) (*dto.EmptyStruct, error) {
	endpoint := fmt.Sprintf(endpoints.CharactersAll, 0)
	return uniapi.Get[dto.EmptyStruct](api, endpoint, nil)
}

// GetAll2 获取总角色简洁信息
func GetAll2(api *uniapi.UniAPI) (*dto.CharactersAll2, error) {
	endpoint := fmt.Sprintf(endpoints.CharactersAll, 2)
	return uniapi.Get[dto.CharactersAll2](api, endpoint, nil)
}

// GetAll5 获取总角色较详细信息
func GetAll5(api *uniapi.UniAPI) (*dto.CharactersAll5, error) {
	endpoint := fmt.Sprintf(endpoints.CharactersAll, 5)
	return uniapi.Get[dto.CharactersAll5](api, endpoint, nil)
}

// GetMain1 获取主要角色 ID 与其乐队 ID
func GetMain1(api *uniapi.UniAPI) (*dto.CharactersMain1, error) {
	endpoint := fmt.Sprintf(endpoints.CharactersMain, 1)
	return uniapi.Get[dto.CharactersMain1](api, endpoint, nil)
}

// GetMain2 获取主要角色简洁信息
func GetMain2(api *uniapi.UniAPI) (*dto.CharactersMain2, error) {
	endpoint := fmt.Sprintf(endpoints.CharactersMain, 2)
	return uniapi.Get[dto.CharactersMain2](api, endpoint, nil)
}

// GetMain3 获取主要角色较详细信息
func GetMain3(api *uniapi.UniAPI) (*dto.CharactersMain3, error) {
	endpoint := fmt.Sprintf(endpoints.CharactersMain, 3)
	return uniapi.Get[dto.CharactersMain3](api, endpoint, nil)
}

// Character 角色
type Character struct {
	Id   int
	Info *dto.CharacterInfo
	api  *uniapi.UniAPI
}

// GetCharacter 获取角色实例
func GetCharacter(api *uniapi.UniAPI, id int) (*Character, error) {
	endpoint := fmt.Sprintf(endpoints.CharactersInfo, id)
	info, err := uniapi.Get[dto.CharacterInfo](api, endpoint, nil)
	if err != nil {
		return nil, err
	}
	return &Character{
		Id:   id,
		Info: info,
		api:  api,
	}, nil
}

func (c *Character) Names() []*string {
	return c.Info.CharacterName
}

// GetIcon 获取角色图标
func (c *Character) GetIcon() (*[]byte, error) {
	name := fmt.Sprintf("chara_icon_%d", c.Id)
	endpoint := fmt.Sprintf(endpoints.IconPng, name)
	return uniapi.Get[[]byte](c.api, endpoint, nil)
}

// GetKVImage 获取角色主视觉图
func (c *Character) GetKVImage() (*[]byte, error) {
	endpoint := fmt.Sprintf(endpoints.CharactersKvImage, "jp", c.Id)
	return uniapi.Get[[]byte](c.api, endpoint, nil)
}
