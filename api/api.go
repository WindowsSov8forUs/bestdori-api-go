package api

import (
	"fmt"
	"io"
	"net/http/cookiejar"
	"time"

	"github.com/go-resty/resty/v2"
	"golang.org/x/net/publicsuffix"
)

type FilesFormData map[string]struct {
	Name   string
	Reader io.Reader
}

type ResponseError struct {
	Response *resty.Response
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("response error: %v", e.Response)
}

type ResponseStatusError ResponseError

func (e *ResponseStatusError) Error() string {
	return fmt.Sprintf("unexpected status code: %d", e.Response.StatusCode())
}

func (e *ResponseStatusError) StatusCode() int {
	return e.Response.StatusCode()
}

func RaiseForStatus(resp *resty.Response) error {
	if resp.IsError() {
		return &ResponseStatusError{Response: resp}
	}
	return nil
}

type NoContentTypeError ResponseError

func (e *NoContentTypeError) Error() string {
	return "got unexpected empty Content-Type"
}

type API struct {
	baseURL string
	client  *resty.Client
}

func NewAPI(url, proxyURL string, timeout int) *API {
	client := resty.New().
		SetBaseURL(url).
		SetProxy(proxyURL).
		SetTimeout(time.Duration(timeout) * time.Second).
		SetRetryCount(5)

	// 启用 Cookie Jar
	jar, _ := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	client.SetCookieJar(jar)

	return &API{
		baseURL: url,
		client:  client,
	}
}

func (api *API) OnAfterResponse(f resty.ResponseMiddleware) {
	api.client.OnAfterResponse(f)
}

func (api *API) OnBeforeRequest(f resty.RequestMiddleware) {
	api.client.OnBeforeRequest(f)
}

func (api *API) ContentTypeMiddleware() resty.ResponseMiddleware {
	return func(c *resty.Client, r *resty.Response) error {
		contentType := r.Header().Get("Content-Type")
		if contentType == "" {
			return &NoContentTypeError{Response: r}
		}
		return nil
	}
}

func (api *API) buildRequest(params map[string]string, data any, files FilesFormData) (*resty.Request, error) {
	req := api.client.R()

	// 设置请求参数
	if params != nil {
		req.SetQueryParams(params)
	}

	// 设置请求体
	if data != nil {
		req.SetBody(data)
	}

	// 设置文件表单
	for field, file := range files {
		req.SetFileReader(field, file.Name, file.Reader)
	}

	return req, nil
}

func (api *API) get(endpoint string, params map[string]string, value any) (*resty.Response, error) {
	req, err := api.buildRequest(params, nil, nil)
	if err != nil {
		return nil, err
	}

	if _, ok := value.(*[]byte); ok {
		response, err := req.Get(endpoint)
		if err != nil {
			return nil, err
		}
		return response, nil
	} else {
		response, err := req.SetResult(&value).Get(endpoint)
		if err != nil {
			return nil, err
		}
		return response, nil
	}
}

func (api *API) post(endpoint string, data any, files FilesFormData, value any) (*resty.Response, error) {
	req, err := api.buildRequest(nil, data, files)
	if err != nil {
		return nil, err
	}

	response, err := req.SetResult(&value).Post(endpoint)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func Get[T any](api *API, endpoint string, params map[string]string) (*T, error) {
	var result T
	resp, err := api.get(endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	if _, ok := any(&result).(*[]byte); ok {
		*any(&result).(*[]byte) = resp.Body()
		return &result, nil
	}
	return &result, nil
}

func Post[T any](api *API, endpoint string, data any, files FilesFormData) (*T, error) {
	var result T
	_, err := api.post(endpoint, data, files, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
