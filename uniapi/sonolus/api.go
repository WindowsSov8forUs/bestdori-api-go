package sonolus

import (
	"fmt"

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
	return fmt.Sprintf("code: %s, description: %s, detail: %s", e.Code, e.Description, e.Detail)
}

func (e *SonolusResponseError) Error() string {
	return fmt.Sprintf("sonolus ayachan server responsed an error: `%s`", e.errorInfo())
}

func NewSonolusAPI(proxyURL string, timeout int) *uniapi.UniAPI {
	api := uniapi.NewAPI("https://sonolus.ayachan.fun", proxyURL, timeout)
	api.OnAfterResponse(onAfterResponse)
	api.OnAfterResponse(api.ContentTypeMiddleware())
	return api
}
