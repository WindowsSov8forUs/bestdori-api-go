package niconi

import (
	"github.com/WindowsSov8forUs/bestdori-api-go/api"
	"github.com/go-resty/resty/v2"
)

func onAfterResponse(client *resty.Client, response *resty.Response) error {
	return api.RaiseForStatus(response)
}
