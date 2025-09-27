package uniapi

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

type UniAPI struct {
	baseURL string
	client  *resty.Client
}

func NewAPI(url, proxyURL string, timeout int) *UniAPI {
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

	return &UniAPI{
		baseURL: url,
		client:  client,
	}
}

func (api *UniAPI) OnAfterResponse(f resty.ResponseMiddleware) {
	api.client.OnAfterResponse(f)
}

func (api *UniAPI) OnBeforeRequest(f resty.RequestMiddleware) {
	api.client.OnBeforeRequest(f)
}

func (api *UniAPI) ContentTypeMiddleware() resty.ResponseMiddleware {
	return func(c *resty.Client, r *resty.Response) error {
		contentType := r.Header().Get("Content-Type")
		if contentType == "" {
			return &NoContentTypeError{Response: r}
		}
		return nil
	}
}

func parseParams(params map[string]any) map[string]string {
	result := make(map[string]string, len(params))
	for k, v := range params {
		if v == nil {
			continue
		}

		switch val := v.(type) {
		case string:
			params[k] = val
		case int, int8, int16, int32, int64:
			params[k] = fmt.Sprintf("%d", val)
		case uint, uint8, uint16, uint32, uint64:
			params[k] = fmt.Sprintf("%d", val)
		case float32, float64:
			params[k] = fmt.Sprintf("%g", val)
		case bool:
			params[k] = fmt.Sprintf("%t", val)
		case fmt.Stringer:
			params[k] = val.String()
		default:
			params[k] = fmt.Sprintf("%v", val)
		}
	}
	return result
}

func (api *UniAPI) buildRequest(params map[string]any, data any, files FilesFormData) (*resty.Request, error) {
	req := api.client.R()

	// 设置请求参数
	if params != nil {
		req.SetQueryParams(parseParams(params))
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

func (api *UniAPI) get(endpoint string, params map[string]any, value any) (*resty.Response, error) {
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

func (api *UniAPI) post(endpoint string, data any, files FilesFormData, value any) (*resty.Response, error) {
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

func Get[T any](api *UniAPI, endpoint string, params map[string]any) (*T, error) {
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

func Post[T any](api *UniAPI, endpoint string, data any, files FilesFormData) (*T, error) {
	var result T
	_, err := api.post(endpoint, data, files, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
