package ayachan

import (
	"encoding/json"

	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
	"github.com/go-resty/resty/v2"
)

type AyachanAPIResponse struct {
	Error string `json:"error,omitempty"`
}

type AyachanResponseError struct {
	*uniapi.ResponseError
	ErrorInfo string
}

func (e *AyachanResponseError) Error() string {
	if e.ErrorInfo == "" {
		return "ayachan responsed an error"
	}
	return "ayachan responsed an error: " + e.ErrorInfo
}

func OnAfterResponseAyachan(client *resty.Client, response *resty.Response) error {
	// 处理异常响应
	if response.StatusCode() == 400 || response.StatusCode() == 500 {
		// 解析通用响应体
		var resp AyachanAPIResponse
		if err := json.Unmarshal(response.Body(), &resp); err != nil {
			return err
		}
		return &AyachanResponseError{
			ResponseError: &uniapi.ResponseError{Response: response},
			ErrorInfo:     resp.Error,
		}
	}
	return uniapi.RaiseForStatus(response)
}
