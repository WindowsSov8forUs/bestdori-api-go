package bestdori

import (
	"fmt"

	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

type BestdoriAPIResponse struct {
	Result *bool  `json:"result,omitempty"`
	Code   string `json:"code,omitempty"`
}

type RequestFiledError struct {
	*uniapi.ResponseError
	Code string
}

type AssetsNotExistError uniapi.ResponseError

func (e *AssetsNotExistError) Error() string {
	return fmt.Sprintf("assets or res `%s` not exist", e.Response.Request.URL)
}

func (e *RequestFiledError) Error() string {
	if e.Code == "" {
		return "request failed"
	}
	return fmt.Sprintf("request failed with code `%s`", e.Code)
}

func NewBestdoriAPI(proxyURL string, timeout int) *uniapi.UniAPI {
	api := uniapi.NewAPI("https://bestdori.com", proxyURL, timeout)
	api.OnBeforeRequest(onBeforeRequest)
	api.OnAfterResponse(api.ContentTypeMiddleware())
	api.OnAfterResponse(onAfterResponse)
	return api
}
