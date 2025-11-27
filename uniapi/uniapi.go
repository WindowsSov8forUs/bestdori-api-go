package uniapi

import (
	"fmt"
	"io"
	"net/http/cookiejar"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-resty/resty/v2"
	"golang.org/x/net/publicsuffix"
)

var logger resty.Logger

func RegisterLogger(l resty.Logger) {
	logger = l
}

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
	baseURL    string
	jar        *cookiejar.Jar
	client     *resty.Client
	proxyURL   string
	timeout    time.Duration
	log        resty.Logger
	recreateMu sync.Mutex
	invalid    atomic.Bool
}

func NewAPI(url, proxyURL string, timeout int) *UniAPI {
	if timeout <= 0 {
		timeout = 30
	}
	api := &UniAPI{
		baseURL:  url,
		proxyURL: proxyURL,
		timeout:  time.Duration(timeout) * time.Second,
		log:      logger,
	}
	jar, _ := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	api.jar = jar
	client := resty.New().
		SetBaseURL(url).
		SetRetryCount(5).
		SetTimeout(api.timeout)
	if proxyURL != "" {
		client.SetProxy(proxyURL)
	}
	if logger != nil {
		client.SetLogger(logger).
			SetDebug(true)
	} else {
		client.SetDebug(false)
	}
	client.SetCookieJar(api.jar)
	api.client = client
	return api
}

// SetCookies 从响应中设置 Cookie
func (api *UniAPI) SetCookies(resp *resty.Response) {
	if resp != nil {
		api.client.SetCookies(resp.Cookies())
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
			result[k] = val
		case int:
			result[k] = strconv.Itoa(val)
		case int8:
			result[k] = strconv.FormatInt(int64(val), 10)
		case int16:
			result[k] = strconv.FormatInt(int64(val), 10)
		case int32:
			result[k] = strconv.FormatInt(int64(val), 10)
		case int64:
			result[k] = strconv.FormatInt(val, 10)
		case uint:
			result[k] = strconv.FormatUint(uint64(val), 10)
		case uint8:
			result[k] = strconv.FormatUint(uint64(val), 10)
		case uint16:
			result[k] = strconv.FormatUint(uint64(val), 10)
		case uint32:
			result[k] = strconv.FormatUint(uint64(val), 10)
		case uint64:
			result[k] = strconv.FormatUint(val, 10)
		case float32:
			result[k] = strconv.FormatFloat(float64(val), 'g', -1, 32)
		case float64:
			result[k] = strconv.FormatFloat(val, 'g', -1, 64)
		case bool:
			result[k] = strconv.FormatBool(val)
		case fmt.Stringer:
			result[k] = val.String()
		default:
			result[k] = fmt.Sprintf("%v", val)
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
	const maxAttempts = 5
	for attempt := range maxAttempts {
		if attempt > 0 {
			time.Sleep(time.Duration(1<<attempt) * time.Second)
		}
		if api.invalid.Load() {
			api.recreateClient()
		}
		req, err := api.buildRequest(params, nil, nil)
		if err != nil {
			return nil, err
		}
		var resp *resty.Response
		var execErr error
		if _, ok := value.(*[]byte); ok {
			resp, execErr = req.Get(endpoint)
		} else {
			resp, execErr = req.SetResult(value).Get(endpoint)
		}
		if execErr == nil || !api.shouldRecreate(execErr) || attempt == maxAttempts-1 {
			return resp, execErr
		}
		api.invalid.Store(true)
	}
	return nil, fmt.Errorf("get failed after %d attempts with client recreation", maxAttempts)
}

func (api *UniAPI) post(endpoint string, data any, files FilesFormData, value any) (*resty.Response, error) {
	const maxAttempts = 2
	for attempt := range maxAttempts {
		if api.invalid.Load() {
			api.recreateClient()
		}
		req, err := api.buildRequest(nil, data, files)
		if err != nil {
			return nil, err
		}
		var resp *resty.Response
		var execErr error
		hasFiles := len(files) > 0
		if _, ok := value.(*resty.Response); ok {
			resp, execErr = req.Post(endpoint)
		} else {
			resp, execErr = req.SetResult(value).Post(endpoint)
		}
		if execErr == nil || !api.shouldRecreate(execErr) || hasFiles || attempt == maxAttempts-1 {
			return resp, execErr
		}
		api.invalid.Store(true)
	}
	return nil, fmt.Errorf("post failed after %d attempts with client recreation", maxAttempts)
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

func (api *UniAPI) shouldRecreate(err error) bool {
	if err == nil {
		return false
	}
	s := strings.ToLower(err.Error())
	return strings.Contains(s, "connection") ||
		strings.Contains(s, "reset by peer") ||
		strings.Contains(s, "use of closed") ||
		strings.Contains(s, "broken pipe") ||
		strings.Contains(s, "eof") ||
		strings.Contains(s, "timeout") ||
		strings.Contains(s, "deadline")
}

func (api *UniAPI) recreateClient() {
	api.recreateMu.Lock()
	defer api.recreateMu.Unlock()
	if !api.invalid.Load() {
		return
	}
	newClient := resty.New().
		SetBaseURL(api.baseURL).
		SetRetryCount(0).
		SetTimeout(api.timeout)
	if api.proxyURL != "" {
		newClient.SetProxy(api.proxyURL)
	}
	if api.log != nil {
		newClient.SetLogger(api.log).
			SetDebug(true)
	} else {
		newClient.SetDebug(false)
	}
	newClient.SetCookieJar(api.jar)
	api.client = newClient
	api.invalid.Store(false)
}
