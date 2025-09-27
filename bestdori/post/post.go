package post

import (
	"fmt"

	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/charts"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

type Order string

const (
	OrderTimeDesc Order = "TIME_DESC"
	OrderTimeAsc  Order = "TIME_ASC"
)

// GetList 搜索社区帖子
func GetList(
	api *uniapi.UniAPI,
	search *string,
	following *bool,
	categoryName, categoryId *string,
	tags *[]dto.PostTag,
	username *string,
	order Order,
	limit int,
	offset int,
) (*dto.PostList, error) {
	data := map[string]any{
		"search":       search,
		"following":    following,
		"categoryName": categoryName,
		"categoryId":   categoryId,
		"tags":         tags,
		"username":     username,
		"order":        order,
		"limit":        limit,
		"offset":       offset,
	}
	return uniapi.Post[dto.PostList](api, endpoints.PostList, data, nil)
}

// SearchTags 搜索现有标签
func SearchTags(api *uniapi.UniAPI, typ, data string, fuzzy bool) (*[]dto.TagResultTag, error) {
	params := map[string]any{
		"type":  typ,
		"data":  data,
		"fuzzy": fuzzy,
	}
	result, err := uniapi.Get[dto.TagResult](api, endpoints.PostTag, params)
	if err != nil {
		return nil, err
	}
	return &result.Tags, nil
}

// CreatePost 发表帖子
func CreatePost(
	api *uniapi.UniAPI,
	Artists *string,
	CategoryId, CategoryName string,
	Chart *[]map[string]any,
	Content []dto.PostContent,
	Diff *dto.ChartDifficulty,
	Level *int,
	Song *dto.PostSong,
	Tags *[]dto.PostTag,
	Title *string,
) (int, error) {
	data := map[string]any{
		"categoryId":   CategoryId,
		"categoryName": CategoryName,
		"content":      Content,
	}

	// 可选字段
	if Artists != nil {
		data["artists"] = Artists
	}
	if Chart != nil {
		data["chart"] = Chart
	}
	if Diff != nil {
		data["diff"] = Diff
	}
	if Level != nil {
		data["level"] = Level
	}
	if Song != nil {
		data["song"] = Song
	}
	if Tags != nil {
		data["tags"] = Tags
	}
	if Title != nil {
		data["title"] = Title
	}

	type resp struct {
		Id int `json:"id"`
	}
	result, err := uniapi.Post[resp](api, endpoints.PostPost, data, nil)
	if err != nil {
		return 0, err
	}
	return result.Id, nil
}

// FindPost 查询帖子顺序
func FindPost(api *uniapi.UniAPI, categoryName, categoryId string, id int) (int, error) {
	params := map[string]any{
		"categoryName": categoryName,
		"categoryId":   categoryId,
		"id":           id,
	}
	type resp struct {
		Position int `json:"position"`
	}
	result, err := uniapi.Get[resp](api, endpoints.PostFind, params)
	if err != nil {
		return 0, err
	}
	return result.Position, nil
}

// Post 帖子
type Post struct {
	Id      int
	Info    *dto.PostInfo
	bdApi   *uniapi.UniAPI
	nicoApi *uniapi.UniAPI
}

// GetPost 获取帖子实例
func GetPost(bdApi, nicoApi *uniapi.UniAPI, id int) (*Post, error) {
	params := map[string]any{"id": id}
	result, err := uniapi.Get[dto.PostDetail](bdApi, endpoints.PostDetails, params)
	if err != nil {
		return nil, err
	}
	return &Post{
		Id:      id,
		Info:    &result.Post,
		bdApi:   bdApi,
		nicoApi: nicoApi,
	}, nil
}

// GetBasic 获取帖子简略信息
func (p *Post) GetBasic() (*dto.PostBasic, error) {
	params := map[string]any{"id": p.Id}
	return uniapi.Get[dto.PostBasic](p.bdApi, endpoints.PostBasic, params)
}

// GetChart 获取帖子中的谱面信息
func (p *Post) GetChart() (*charts.Chart, error) {
	if p.Info.Chart == nil {
		return nil, &bestdori.NotExistError{Target: fmt.Sprintf("chart of post %d", p.Id)}
	}
	return charts.UnmarshalSlice(*p.Info.Chart)
}

// Content 获取帖子内容
func (p *Post) Content() string {
	result := ""
	for _, item := range p.Info.Content {
		switch item.Type {
		case dto.PostContentTypeText, dto.PostContentTypeLink:
			result += item.Data
		case dto.PostContentTypeEmoji:
			result += fmt.Sprintf(":%s:", item.Data)
		case dto.PostContentTypeBr:
			result += "\n"
		}
	}
	return result
}

// GetSong 获取谱面歌曲信息
func (p *Post) GetSong() (*[]byte, *[]byte, error) {
	if p.Info.Song == nil {
		return nil, nil, &bestdori.NotExistError{Target: fmt.Sprintf("song of post %d", p.Id)}
	}
	song := p.Info.Song

	var audio, cover *[]byte
	switch song.Type {
	case dto.PostSongTypeCustom:
		// 自定义歌曲
		if song.Audio == "" {
			audio = nil
		} else {
			endpoints := bestdori.RemoveURLPrefix(song.Audio)
			audio, _ = uniapi.Get[[]byte](p.nicoApi, endpoints, nil)
		}
		if song.Cover == "" {
			cover = nil
		} else {
			endpoints := bestdori.RemoveURLPrefix(song.Cover)
			cover, _ = uniapi.Get[[]byte](p.nicoApi, endpoints, nil)
		}
	case dto.PostSongTypeBandori:
		// BanG Dream! 歌曲
		endpoint := fmt.Sprintf(endpoints.SongsInfo, song.Id)
		info, err := uniapi.Get[dto.SongInfo](p.nicoApi, endpoint, nil)
		if err != nil {
			return nil, nil, err
		}

		// 获取歌曲所在服务器
		var server dto.ServerName
		publishedAt := info.PublishedAt
		if publishedAt[0] != nil {
			server = dto.ServerNameJP
		} else if publishedAt[1] != nil {
			server = dto.ServerNameEN
		} else if publishedAt[2] != nil {
			server = dto.ServerNameTW
		} else if publishedAt[3] != nil {
			server = dto.ServerNameCN
		} else if publishedAt[4] != nil {
			server = dto.ServerNameKR
		} else {
			return nil, nil, &bestdori.NotExistError{
				Target: fmt.Sprintf("server of song %d in post %d", song.Id, p.Id),
			}
		}

		// 获取音频
		audio, _ = uniapi.Get[[]byte](p.nicoApi, fmt.Sprintf(endpoints.SongsSound, server, song.Id, song.Id), nil)

		// 获取封面
		var index int = 0
		quotient, remainder := song.Id/100, song.Id%100
		if remainder == 0 {
			index = song.Id
		} else {
			index = (quotient + 1) * 10
		}
		jacketImage := info.JacketImage
		endpoint = fmt.Sprintf(
			endpoints.SongsMusicJacket,
			server, index, index, jacketImage[len(jacketImage)-1],
		)
		cover, _ = uniapi.Get[[]byte](p.bdApi, endpoint, nil)
	case dto.PostSongTypeLLSIF:
		endpoint := fmt.Sprintf(endpoints.MiscLLSif, 10)
		misc, err := uniapi.Get[dto.LLSifMisc](p.nicoApi, endpoint, nil)
		if err != nil {
			return nil, nil, err
		}
		info, ok := (*misc)[fmt.Sprintf("%d", song.Id)]
		if !ok {
			return nil, nil, &bestdori.NotExistError{
				Target: fmt.Sprintf("song %d in llsif", song.Id),
			}
		}

		// 获取音频
		endpoint = fmt.Sprintf("/%s", info.SoundAsset)
		audio, _ = uniapi.Get[[]byte](p.nicoApi, endpoint, nil)

		// 获取封面
		endpoint = fmt.Sprintf("/%s", info.LiveIconAsset)
		cover, _ = uniapi.Get[[]byte](p.nicoApi, endpoint, nil)
	default:
		return nil, nil, &bestdori.NotExistError{Target: fmt.Sprintf("song type `%s` of post %d", song.Type, p.Id)}
	}

	return audio, cover, nil
}

// GetComments 获取帖子评论
func (p *Post) GetComments(limit, offset int, order Order) (*dto.PostList, error) {
	categoryName := "POST_COMMENT"
	categoryId := fmt.Sprintf("%d", p.Id)

	return GetList(
		p.bdApi,
		nil, nil,
		&categoryName,
		&categoryId,
		nil, nil,
		order,
		limit,
		offset,
	)
}

// Comment 评论帖子
func (p *Post) Comment(content []dto.PostContent) (int, error) {
	return CreatePost(
		p.bdApi,
		nil,
		fmt.Sprintf("%d", p.Id),
		"POST_COMMENT",
		nil,
		content,
		nil, nil, nil, nil, nil,
	)
}

// Like 喜欢 / 取消喜欢帖子
func (p *Post) Like(value bool) error {
	data := map[string]any{
		"id":    p.Id,
		"value": value,
	}
	_, err := uniapi.Post[any](p.bdApi, endpoints.PostLike, data, nil)
	return err
}
