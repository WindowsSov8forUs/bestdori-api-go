package ayachan

import (
	"fmt"

	"github.com/WindowsSov8forUs/bestdori-api-go/api"
)

type AyachanAPIResponse struct {
	Error string `json:"error,omitempty"`
}

type AyachanResponseError struct {
	*api.ResponseError
	ErrorInfo string
}

func (e *AyachanResponseError) Error() string {
	if e.ErrorInfo == "" {
		return "ayachan responsed an error"
	}
	return fmt.Sprintf("ayachan responsed an error: %s", e.ErrorInfo)
}

func NewAyachanAPI(proxyURL string, timeout int) *api.API {
	api := api.NewAPI("https://api.ayachan.fun", proxyURL, timeout)
	api.OnAfterResponse(onAfterResponse)
	api.OnAfterResponse(api.ContentTypeMiddleware())
	return api
}
