package sonolus

import (
	"encoding/json"

	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
	"github.com/go-resty/resty/v2"
)

func onAfterResponse(client *resty.Client, response *resty.Response) error {
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
