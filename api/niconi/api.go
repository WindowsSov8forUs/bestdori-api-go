package niconi

import (
	"github.com/WindowsSov8forUs/bestdori-api-go/api"
)

func NewNiconiAPI(proxyURL string, timeout int) *api.API {
	api := api.NewAPI("https://card.niconi.co.ni", proxyURL, timeout)
	api.OnAfterResponse(onAfterResponse)
	api.OnAfterResponse(api.ContentTypeMiddleware())
	return api
}
