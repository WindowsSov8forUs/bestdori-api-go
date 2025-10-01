package bestdori

import (
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
	"github.com/go-resty/resty/v2"
)

func OnAfterResponseNiconi(client *resty.Client, response *resty.Response) error {
	return uniapi.RaiseForStatus(response)
}
