package bestdori

import (
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
	return "assets or res `" + e.Response.Request.URL + "` not exist"
}

func (e *RequestFiledError) Error() string {
	if e.Code == "" {
		return "request failed"
	}
	return "request failed with code `" + e.Code + "`"
}

func NewBestdoriAPI(proxyURL string, timeout int) *uniapi.UniAPI {
	api := uniapi.NewAPI("https://bestdori.com", proxyURL, timeout)
	api.OnBeforeRequest(onBeforeRequest)
	api.OnAfterResponse(api.ContentTypeMiddleware())
	api.OnAfterResponse(onAfterResponse)
	return api
}
