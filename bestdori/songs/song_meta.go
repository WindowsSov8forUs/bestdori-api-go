package songs

import (
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// GetSongMeta 获取总歌曲 Meta 信息
func GetSongMeta(api *uniapi.UniAPI, index int) (*dto.SongsMetaAll, error) {
	endpoint := endpoints.MetaAll(index)
	return uniapi.Get[dto.SongsMetaAll](api, endpoint, nil)
}
