package upload

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

type uploadResponse struct {
	Hash string `json:"hash"`
}

type statusResponse struct {
	Status string `json:"status"`
}

// Upload 上传
type Upload struct {
	ver    int
	hash   string
	size   int64
	name   string
	reader io.Reader
	api    *uniapi.UniAPI
}

// NewUpload 创建上传对象
func NewUpload(api *uniapi.UniAPI, name string, reader io.Reader) *Upload {
	// 读取文件内容，计算hash和size
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil
	}

	// 计算 SHA-1 哈希值与文件大小
	hashBytes := sha1.Sum(data)
	hash := hex.EncodeToString(hashBytes[:])
	size := int64(len(data))

	return &Upload{
		ver:    3,
		name:   name,
		hash:   hash,
		size:   size,
		reader: bytes.NewReader(data),
		api:    api,
	}
}

// NewUploadFromPath 从路径创建上传对象
func NewUploadFromPath(api *uniapi.UniAPI, path string) (*Upload, error) {
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return NewUpload(api, filepath.Base(path), file), nil
}

// Upload 上传文件
func (u *Upload) Upload() (string, error) {
	// 发送预上传请求
	data := map[string]any{
		"ver":  u.ver,
		"hash": u.hash,
		"size": u.size,
	}
	if _, err := uniapi.Post[any](u.api, endpoints.UploadPrepare(), data, nil); err != nil {
		if reqFailedErr, ok := err.(*bestdori.RequestFailedError); ok {
			if reqFailedErr.Code == "ALREADY_UPLOADED" {
				return u.hash, nil
			}
		}
		return "", err
	}

	// 发送上传请求
	files := uniapi.FilesFormData{
		"file": {
			Name:   u.name,
			Reader: u.reader,
		},
	}
	uploadResp, err := uniapi.Post[uploadResponse](u.api, endpoints.UploadUpload(), data, files)
	if err != nil {
		return "", err
	}

	hashGet := uploadResp.Hash
	// 重复查询至多 5 次上传状态
	endpoint := endpoints.UploadStatus(hashGet)
	for i := 0; i < 5; i++ {
		if statusResp, err := uniapi.Get[statusResponse](u.api, endpoint, nil); err != nil {
			return "", err
		} else {
			if statusResp.Status == "available" {
				return hashGet, nil
			}
		}
	}

	return "", fmt.Errorf("upload file " + u.name + " timeout")
}
