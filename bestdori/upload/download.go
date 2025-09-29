package upload

import (
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// Download 下载已上传到 Bestdori 的文件
func Download(api *uniapi.UniAPI, hash string) (*[]byte, error) {
	endpoint := endpoints.UploadFile(hash)
	return uniapi.Get[[]byte](api, endpoint, nil)
}
