package bestdori

import (
	"encoding/json"
	"strings"

	"github.com/WindowsSov8forUs/bestdori-api-go/api"
	"github.com/go-resty/resty/v2"
)

func onBeforeRequest(client *resty.Client, request *resty.Request) error {
	// 设置请求头
	if !strings.HasSuffix(request.URL, "/api/upload") {
		request.SetHeader("Content-Type", "application/json;charset=UTF-8")
	}
	return nil
}

func onAfterResponse(client *resty.Client, response *resty.Response) error {
	// 处理异常响应
	if strings.Contains(response.Request.URL, "/api/") {
		// 检查 Content-Type
		contentType := response.Header().Get("Content-Type")
		if !strings.Contains(contentType, "application/json") {
			return api.RaiseForStatus(response)
		}

		// 解析通用响应体
		var resp BestdoriAPIResponse
		if err := json.Unmarshal(response.Body(), &resp); err != nil {
			return err
		}
		if resp.Result == nil {
			if response.IsError() {
				return &api.ResponseStatusError{Response: response}
			}
			return nil
		} else if !*resp.Result {
			code := resp.Code
			return &RequestFiledError{
				ResponseError: &api.ResponseError{Response: response},
				Code:          code,
			}
		}
		return api.RaiseForStatus(response)
	} else if strings.Contains(response.Request.URL, "/assets/") || strings.Contains(response.Request.URL, "/res/") {
		// 检查 Content-Type
		if strings.Contains(response.Header().Get("Content-Type"), "text/html") {
			return &AssetsNotExistError{Response: response}
		}
		return api.RaiseForStatus(response)
	}
	return api.RaiseForStatus(response)
}
