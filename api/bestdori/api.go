package bestdori

import (
	"fmt"

	"github.com/WindowsSov8forUs/bestdori-api-go/api"
)

type BestdoriAPIResponse struct {
	Result *bool  `json:"result,omitempty"`
	Code   string `json:"code,omitempty"`
}

type RequestFiledError struct {
	*api.ResponseError
	Code string
}

type AssetsNotExistError api.ResponseError

func (e *AssetsNotExistError) Error() string {
	return fmt.Sprintf("assets or res `%s` not exist", e.Response.Request.URL)
}

func (e *RequestFiledError) Error() string {
	if e.Code == "" {
		return "request failed"
	}
	return fmt.Sprintf("request failed with code `%s`", e.Code)
}

func NewBestdoriAPI(proxyURL string, timeout int) *api.API {
	api := api.NewAPI("https://bestdori.com", proxyURL, timeout)
	api.OnBeforeRequest(onBeforeRequest)
	api.OnAfterResponse(onAfterResponse)
	return api
}
