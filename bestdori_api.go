// Package bestdoriapi provides a Go SDK for accessing various BanG Dream related APIs
package bestdoriapi

import (
	"github.com/WindowsSov8forUs/bestdori-api-go/ayachan"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
	"github.com/go-resty/resty/v2"
)

// NewBestdoriAPI creates a new Bestdori API client
func NewBestdoriAPI(proxyURL string, timeout int) *uniapi.UniAPI {
	api := uniapi.NewAPI("https://bestdori.com", proxyURL, timeout)
	api.OnBeforeRequest(bestdori.OnBeforeRequestBestdori)
	api.OnAfterResponse(api.ContentTypeMiddleware())
	api.OnAfterResponse(bestdori.OnAfterResponseBestdori)
	return api
}

// NewNiconiAPI creates a new Niconi API client
func NewNiconiAPI(proxyURL string, timeout int) *uniapi.UniAPI {
	api := uniapi.NewAPI("https://card.niconi.co.ni", proxyURL, timeout)
	api.OnAfterResponse(bestdori.OnAfterResponseNiconi)
	api.OnAfterResponse(api.ContentTypeMiddleware())
	return api
}

// NewAyachanAPI creates a new Ayachan API client
func NewAyachanAPI(proxyURL string, timeout int) *uniapi.UniAPI {
	api := uniapi.NewAPI("https://api.ayachan.fun/v2", proxyURL, timeout)
	api.OnAfterResponse(ayachan.OnAfterResponseAyachan)
	api.OnAfterResponse(api.ContentTypeMiddleware())
	return api
}

// NewSonolusAPI creates a new Sonolus API client
func NewSonolusAPI(proxyURL string, timeout int) *uniapi.UniAPI {
	api := uniapi.NewAPI("https://sonolus.ayachan.fun/test/sonolus", proxyURL, timeout)
	api.OnAfterResponse(ayachan.OnAfterResponseSonolus)
	api.OnAfterResponse(api.ContentTypeMiddleware())
	return api
}

func RegisterLogger(l resty.Logger) {
	uniapi.RegisterLogger(l)
}
