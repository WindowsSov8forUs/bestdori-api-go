package sonolus

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/WindowsSov8forUs/bestdori-api-go/ayachan/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/ayachan/endpoints"
	bdDto "github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

type levelsPostResponse struct {
	UID int `json:"uid"`
}

// LevelsPost Sonolus 测试谱面上传
func LevelsPost(snlsApi *uniapi.UniAPI, title, bgm string, chart *bdDto.Chart, difficulty int, hidden bool, lifetime int) (int, error) {
	// 打开 BGM 文件
	file, err := os.Open(bgm)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// 将谱面转换为 JSON
	chartJSON, err := json.Marshal(chart)
	if err != nil {
		return 0, err
	}

	// 构建数据
	data := map[string]any{
		"title":      title,
		"chart":      string(chartJSON),
		"difficulty": difficulty,
		"lifetime":   lifetime,
	}
	if hidden {
		data["hidden"] = true
	}

	// 准备文件
	files := uniapi.FilesFormData{
		"bgm": {
			Name:   filepath.Base(bgm),
			Reader: file,
		},
	}

	resp, err := uniapi.Post[levelsPostResponse](snlsApi, endpoints.LevelsPost(), data, files)
	if err != nil {
		return 0, err
	}
	if resp.UID == 0 {
		return 0, fmt.Errorf("got response without `uid` field")
	}
	return resp.UID, nil
}

// LevelsGet Sonolus 测试谱面获取
func LevelsGet(snlsApi *uniapi.UniAPI, uid int) (*bdDto.Chart, error) {
	return uniapi.Get[bdDto.Chart](snlsApi, endpoints.LevelsGet(uid), nil)
}

// Levels Sonolus 测试服谱面信息获取
func Levels(snlsApi *uniapi.UniAPI, uid int) (*dto.Level, error) {
	return uniapi.Get[dto.Level](snlsApi, endpoints.LevelsInfo(uid), nil)
}
