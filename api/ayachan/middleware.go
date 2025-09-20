package ayachan

import (
	"encoding/json"

	"github.com/WindowsSov8forUs/bestdori-api-go/api"
	"github.com/go-resty/resty/v2"
)

func onAfterResponse(client *resty.Client, response *resty.Response) error {
	// 处理异常响应
	if response.StatusCode() == 400 || response.StatusCode() == 500 {
		// 解析通用响应体
		var resp AyachanAPIResponse
		if err := json.Unmarshal(response.Body(), &resp); err != nil {
			return err
		}
		return &AyachanResponseError{
			ResponseError: &api.ResponseError{Response: response},
			ErrorInfo:     resp.Error,
		}
	}
	return api.RaiseForStatus(response)
}
