package comics

import (
	"strconv"
	"strings"

	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/post"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// GetAll5 获取总漫画信息
func GetAll5(api *uniapi.UniAPI) (*dto.ComicsAll5, error) {
	endpoint := endpoints.ComicsAll(5)
	return uniapi.Get[dto.ComicsAll5](api, endpoint, nil)
}

// Comic 漫画
type Comic struct {
	Id   int
	Info dto.ComicInfo
	api  *uniapi.UniAPI
}

// GetComic 获取漫画实例
func GetComic(api *uniapi.UniAPI, id int) (*Comic, error) {
	all, err := GetAll5(api)
	if err != nil {
		return nil, err
	}
	info, ok := (*all)[strconv.Itoa(id)]
	if !ok {
		return nil, &bestdori.NotExistError{Target: "comic " + strconv.Itoa(id)}
	}
	return &Comic{
		Id:   id,
		Info: info,
		api:  api,
	}, nil
}

func (c *Comic) Names() []*string {
	return c.Info.Title
}

func (c *Comic) DefaultServer() dto.ServerName {
	publicStartAt := c.Info.PublicStartAt
	switch {
	case publicStartAt[0] != nil:
		return dto.ServerNameJP
	case publicStartAt[1] != nil:
		return dto.ServerNameEN
	case publicStartAt[2] != nil:
		return dto.ServerNameTW
	case publicStartAt[3] != nil:
		return dto.ServerNameCN
	case publicStartAt[4] != nil:
		return dto.ServerNameKR
	default:
		return ""
	}
}

func (c *Comic) Type() string {
	assetBundleName := c.Info.AssetBundleName
	if strings.Contains(assetBundleName, "fourframe") {
		return "fourframe"
	} else {
		return "singleframe"
	}
}

// GetComments 获取漫画评论
func (c *Comic) GetComments(limit, offset int, order post.Order) (*dto.PostList, error) {
	return post.GetList(
		c.api,
		"", false,
		"COMIC_COMMENT",
		strconv.Itoa(c.Id),
		nil, "",
		order,
		limit,
		offset,
	)
}

// GetThumbnail 获取漫画缩略图
func (c *Comic) GetThumbnail(server dto.ServerName) (*[]byte, error) {
	publicStartAt := c.Info.PublicStartAt
	serverId := server.Id()
	if publicStartAt[serverId] == nil {
		return nil, &bestdori.ServerNotAvailableError{
			Target: "comic " + strconv.Itoa(c.Id),
			Server: server,
		}
	}

	assetBundleName := c.Info.AssetBundleName
	endpoint := endpoints.ComicThumbnail(string(server), c.Type(), assetBundleName)
	return uniapi.Get[[]byte](c.api, endpoint, nil)
}

// GetAsset 获取漫画资源图片
func (c *Comic) GetAsset(server dto.ServerName) (*[]byte, error) {
	publicStartAt := c.Info.PublicStartAt
	serverId := server.Id()
	if publicStartAt[serverId] == nil {
		return nil, &bestdori.ServerNotAvailableError{
			Target: "comic " + strconv.Itoa(c.Id),
			Server: server,
		}
	}

	assetBundleName := c.Info.AssetBundleName
	endpoint := endpoints.ComicComic(string(server), c.Type(), assetBundleName)
	return uniapi.Get[[]byte](c.api, endpoint, nil)
}
