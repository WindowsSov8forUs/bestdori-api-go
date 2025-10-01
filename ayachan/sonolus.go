package ayachan

import (
	"encoding/json"

	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
	"github.com/go-resty/resty/v2"
)

type SonolusAPIResponse struct {
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
	Detail      string `json:"detail,omitempty"`
}

type SonolusResponseError struct {
	*uniapi.ResponseError
	Code        string
	Description string
	Detail      string
}

func (e *SonolusResponseError) errorInfo() string {
	return "code: " + e.Code + ", description: " + e.Description + ", detail: " + e.Detail
}

func (e *SonolusResponseError) Error() string {
	return "sonolus ayachan server responsed an error: `" + e.errorInfo() + "`"
}

func OnAfterResponseSonolus(client *resty.Client, response *resty.Response) error {
	// 处理异常响应
	if response.StatusCode() != 200 {
		// 解析通用响应体
		var resp SonolusAPIResponse
		if err := json.Unmarshal(response.Body(), &resp); err != nil {
			return err
		}
		return &SonolusResponseError{
			ResponseError: &uniapi.ResponseError{Response: response},
			Code:          resp.Code,
			Description:   resp.Description,
			Detail:        resp.Detail,
		}
	}
	return uniapi.RaiseForStatus(response)
}
