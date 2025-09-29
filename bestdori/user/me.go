package user

import (
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
	"github.com/go-resty/resty/v2"
)

// Me 当前用户
type Me struct {
	Username string
	Password string
	Me       *dto.MeInfo
	api      *uniapi.UniAPI
}

// Login 登录并获取当前用户实例
func Login(api *uniapi.UniAPI, username, password string) (*Me, error) {
	// 登录用户
	data := map[string]any{
		"username": username,
		"password": password,
	}
	if resp, err := uniapi.Post[resty.Response](api, endpoints.UserLogin(), data, nil); err != nil {
		return nil, err
	} else {
		api.SetCookies(resp)
	}

	// 获取当前用户信息
	if info, err := uniapi.Get[dto.MeInfo](api, endpoints.UserMe(), nil); err != nil {
		return nil, err
	} else {
		return &Me{
			Username: username,
			Password: password,
			Me:       info,
			api:      api,
		}, nil
	}
}

// GetUser 获取当前用户的用户实例
func (m *Me) GetUser() (*User, error) {
	return GetUser(m.api, m.Username)
}

// UpdateInfo 更新当前用户信息
func (m *Me) UpdateInfo(info *dto.UserInfo) error {
	if _, err := uniapi.Post[any](m.api, endpoints.UserInfo(), info, nil); err != nil {
		return err
	}
	return nil
}
