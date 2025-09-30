package sonolus

import (
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
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

func NewSonolusAPI(proxyURL string, timeout int) *uniapi.UniAPI {
	api := uniapi.NewAPI("https://sonolus.ayachan.fun/test/sonolus", proxyURL, timeout)
	api.OnAfterResponse(onAfterResponse)
	api.OnAfterResponse(api.ContentTypeMiddleware())
	return api
}
