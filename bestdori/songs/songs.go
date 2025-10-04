package songs

import (
	"strconv"

	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/charts"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/post"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// GetAll0 获取总歌曲 ID
func GetAll0(api *uniapi.UniAPI) (*dto.EmptyStruct, error) {
	endpoint := endpoints.SongsAll(0)
	return uniapi.Get[dto.EmptyStruct](api, endpoint, nil)
}

// GetAll1 获取总歌曲曲名信息
func GetAll1(api *uniapi.UniAPI) (*dto.SongsAll1, error) {
	endpoint := endpoints.SongsAll(1)
	return uniapi.Get[dto.SongsAll1](api, endpoint, nil)
}

// GetAll5 获取总歌曲简洁信息
func GetAll5(api *uniapi.UniAPI) (*dto.SongsAll5, error) {
	endpoint := endpoints.SongsAll(5)
	return uniapi.Get[dto.SongsAll5](api, endpoint, nil)
}

// GetAll7 获取总歌曲较详细信息
func GetAll7(api *uniapi.UniAPI) (*dto.SongsAll7, error) {
	endpoint := endpoints.SongsAll(7)
	return uniapi.Get[dto.SongsAll7](api, endpoint, nil)
}

// GetAll8 获取总歌曲详细信息
func GetAll8(api *uniapi.UniAPI) (*dto.SongsAll8, error) {
	endpoint := endpoints.SongsAll(8)
	return uniapi.Get[dto.SongsAll8](api, endpoint, nil)
}

// Jacket 歌曲封面
type Jacket struct {
	JacketImage string
	index       int
	server      dto.ServerName
	api         *uniapi.UniAPI
}

func (j *Jacket) Endpoint() string {
	return endpoints.SongsMusicJacket(string(j.server), j.index, j.JacketImage)
}

// Bytes 获取封面图片字节
func (j *Jacket) Bytes() (*[]byte, error) {
	return uniapi.Get[[]byte](j.api, j.Endpoint(), nil)
}

// Song 歌曲
type Song struct {
	Id   int
	Info *dto.SongInfo
	api  *uniapi.UniAPI
}

// GetSong 获取歌曲实例
func GetSong(api *uniapi.UniAPI, id int) (*Song, error) {
	endpoint := endpoints.SongsInfo(id)
	if info, err := uniapi.Get[dto.SongInfo](api, endpoint, nil); err != nil {
		return nil, err
	} else {
		return &Song{
			Id:   id,
			Info: info,
			api:  api,
		}, nil
	}
}

func (s *Song) Names() []*string {
	return s.Info.MusicTitle
}

func (s *Song) DefaultServer() dto.ServerName {
	publishedAt := s.Info.PublishedAt
	switch {
	case publishedAt[0] != nil:
		return dto.ServerNameJP
	case publishedAt[1] != nil:
		return dto.ServerNameEN
	case publishedAt[2] != nil:
		return dto.ServerNameTW
	case publishedAt[3] != nil:
		return dto.ServerNameCN
	case publishedAt[4] != nil:
		return dto.ServerNameKR
	default:
		return ""
	}
}

// GetJacket 获取歌曲封面实例
func (s *Song) GetJacket() []Jacket {
	// 获取数据包序列号
	var index int
	quotient := s.Id / 10
	remainder := s.Id % 10
	if remainder == 0 {
		index = s.Id
	} else {
		index = (quotient + 1) * 10
	}

	jacketImage := s.Info.JacketImage
	var jackets []Jacket = make([]Jacket, 0, len(jacketImage))

	for _, image := range jacketImage {
		jacket := Jacket{
			index:       index,
			JacketImage: image,
			server:      s.DefaultServer(),
			api:         s.api,
		}
		jackets = append(jackets, jacket)
	}
	return jackets
}

// GetChart 获取歌曲谱面
func (s *Song) GetChart(diff dto.ChartDifficultyName) (*charts.Chart, error) {
	if chart, err := charts.GetChart(s.api, s.Id, diff); err == nil {
		return chart, nil
	} else {
		if e, ok := err.(*uniapi.ResponseStatusError); ok {
			if e.StatusCode() == 404 {
				return nil, &bestdori.NotExistError{
					Target: string(diff) + " chart of song " + strconv.Itoa(s.Id),
				}
			}
		}
		return nil, err
	}
}

// GetBGM 获取歌曲音频
func (s *Song) GetBGM() (*[]byte, error) {
	endpoint := endpoints.SongsSound(string(s.DefaultServer()), s.Id)
	return uniapi.Get[[]byte](s.api, endpoint, nil)
}

// GetComments 获取歌曲评论
func (s *Song) GetComments(limit, offset int, order post.Order) (*dto.PostList, error) {
	categoryName := "SONG_COMMENT"
	categoryId := strconv.Itoa(s.Id)

	return post.GetList(
		s.api,
		nil, nil,
		&categoryName,
		&categoryId,
		nil, nil,
		order,
		limit,
		offset,
	)
}
