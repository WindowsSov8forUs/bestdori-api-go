package niconi

import "github.com/WindowsSov8forUs/bestdori-api-go/uniapi"

func NewNiconiAPI(proxyURL string, timeout int) *uniapi.UniAPI {
	api := uniapi.NewAPI("https://card.niconi.co.ni", proxyURL, timeout)
	api.OnAfterResponse(onAfterResponse)
	api.OnAfterResponse(api.ContentTypeMiddleware())
	return api
}
