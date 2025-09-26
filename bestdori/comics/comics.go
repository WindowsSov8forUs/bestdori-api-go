package comics

import (
	"fmt"
	"strings"

	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// GetAll5 获取总漫画信息
func GetAll5(api *uniapi.UniAPI) (*dto.ComicsAll5, error) {
	endpoint := fmt.Sprintf(endpoints.ComicsAll, 5)
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
	info, ok := (*all)[fmt.Sprintf("%d", id)]
	if !ok {
		return nil, &bestdori.NotExistError{Target: fmt.Sprintf("comic %d", id)}
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

func (c *Comic) DefaultServer() dto.Server {
	publicStartAt := c.Info.PublicStartAt
	if publicStartAt[0] != nil {
		return dto.ServerJP
	} else if publicStartAt[1] != nil {
		return dto.ServerEN
	} else if publicStartAt[2] != nil {
		return dto.ServerTW
	} else if publicStartAt[3] != nil {
		return dto.ServerCN
	} else if publicStartAt[4] != nil {
		return dto.ServerKR
	} else {
		return 0
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

// GetThumbnail 获取漫画缩略图
func (c *Comic) GetThumbnail(server dto.ServerName) (*[]byte, error) {
	publicStartAt := c.Info.PublicStartAt
	serverId, err := bestdori.ServerNameToId(server)
	if err != nil {
		return nil, err
	}
	if publicStartAt[serverId] == nil {
		return nil, &bestdori.ServerNotAvailableError{
			Target: fmt.Sprintf("comic %d", c.Id),
			Server: server,
		}
	}

	assetBundleName := c.Info.AssetBundleName
	endpoint := fmt.Sprintf(
		endpoints.ComicThumbnail,
		server, c.Type(), assetBundleName, assetBundleName,
	)
	return uniapi.Get[[]byte](c.api, endpoint, nil)
}

// GetAsset 获取漫画资源图片
func (c *Comic) GetAsset(server dto.ServerName) (*[]byte, error) {
	publicStartAt := c.Info.PublicStartAt
	serverId, err := bestdori.ServerNameToId(server)
	if err != nil {
		return nil, err
	}
	if publicStartAt[serverId] == nil {
		return nil, &bestdori.ServerNotAvailableError{
			Target: fmt.Sprintf("comic %d", c.Id),
			Server: server,
		}
	}

	assetBundleName := c.Info.AssetBundleName
	endpoint := fmt.Sprintf(
		endpoints.ComicComic,
		server, c.Type(), assetBundleName, assetBundleName,
	)
	return uniapi.Get[[]byte](c.api, endpoint, nil)
}
