package user

import (
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/post"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// User 用户
type User struct {
	Username string
	Info     *dto.UserInfo
	api      *uniapi.UniAPI
}

// GetUser 获取用户实例
func GetUser(api *uniapi.UniAPI, username string) (*User, error) {
	params := map[string]any{
		"username": username,
	}
	if info, err := uniapi.Get[dto.UserInfo](api, endpoints.UserInfo(), params); err != nil {
		return nil, err
	} else {
		return &User{
			Username: username,
			Info:     info,
			api:      api,
		}, nil
	}
}

// GetPosts 获取用户的帖子列表
func (u *User) GetPosts(limit, offset int, order post.Order) (*dto.PostList, error) {
	return post.GetList(
		u.api,
		nil, nil, nil, nil, nil,
		&u.Username,
		order,
		limit,
		offset,
	)
}

// GetCharts 获取用户的谱面列表
func (u *User) GetCharts(limit, offset int, order post.Order) (*dto.PostList, error) {
	categoryName := "SELF_POST"
	categoryId := "chart"

	return post.GetList(
		u.api,
		nil, nil,
		&categoryName,
		&categoryId,
		nil,
		&u.Username,
		order,
		limit,
		offset,
	)
}

// GetTexts 获取用户的文本帖子
func (u *User) GetTexts(limit, offset int, order post.Order) (*dto.PostList, error) {
	categoryName := "SELF_POST"
	categoryId := "text"

	return post.GetList(
		u.api,
		nil, nil,
		&categoryName,
		&categoryId,
		nil,
		&u.Username,
		order,
		limit,
		offset,
	)
}

// GetStories 获取用户的故事帖子
func (u *User) GetStories(limit, offset int, order post.Order) (*dto.PostList, error) {
	categoryName := "SELF_POST"
	categoryId := "story"

	return post.GetList(
		u.api,
		nil, nil,
		&categoryName,
		&categoryId,
		nil,
		&u.Username,
		order,
		limit,
		offset,
	)
}
