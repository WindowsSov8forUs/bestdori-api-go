package bestdori

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
	"github.com/go-resty/resty/v2"
)

type NotExistError struct {
	uniapi.ResponseError
	Target string
}

func (e *NotExistError) Error() string {
	return e.Target + " does not exist"
}

type ServerNotAvailableError struct {
	uniapi.ResponseError
	Target string
	Server dto.ServerName
}

func (e *ServerNotAvailableError) Error() string {
	return e.Target + " is not available on server " + string(e.Server)
}

func ServerNameToId(serverName dto.ServerName) (dto.Server, error) {
	switch serverName {
	case dto.ServerNameJP:
		return dto.ServerJP, nil
	case dto.ServerNameEN:
		return dto.ServerEN, nil
	case dto.ServerNameTW:
		return dto.ServerTW, nil
	case dto.ServerNameCN:
		return dto.ServerCN, nil
	case dto.ServerNameKR:
		return dto.ServerKR, nil
	default:
		return 0, fmt.Errorf("unknown server name: " + string(serverName))
	}
}

func ServerIdToName(serverId dto.Server) (dto.ServerName, error) {
	switch serverId {
	case dto.ServerJP:
		return dto.ServerNameJP, nil
	case dto.ServerEN:
		return dto.ServerNameEN, nil
	case dto.ServerTW:
		return dto.ServerNameTW, nil
	case dto.ServerCN:
		return dto.ServerNameCN, nil
	case dto.ServerKR:
		return dto.ServerNameKR, nil
	default:
		return "", fmt.Errorf("unknown server id: " + strconv.Itoa(int(serverId)))
	}
}

func RemoveURLPrefix(url string) string {
	if strings.HasPrefix(url, "https://bestdori.com") {
		return strings.TrimPrefix(url, "https://bestdori.com")
	}
	return url
}

type BestdoriAPIResponse struct {
	Result *bool  `json:"result,omitempty"`
	Code   string `json:"code,omitempty"`
}

type RequestFailedError struct {
	*uniapi.ResponseError
	Code string
}

type AssetsNotExistError uniapi.ResponseError

func (e *AssetsNotExistError) Error() string {
	return "assets or res `" + e.Response.Request.URL + "` not exist"
}

func (e *RequestFailedError) Error() string {
	if e.Code == "" {
		return "request failed"
	}
	return "request failed with code `" + e.Code + "`"
}

func OnBeforeRequestBestdori(client *resty.Client, request *resty.Request) error {
	// 设置请求头
	if !strings.HasSuffix(request.URL, "/api/upload") {
		request.SetHeader("Content-Type", "application/json;charset=UTF-8")
	}
	return nil
}

func OnAfterResponseBestdori(client *resty.Client, response *resty.Response) error {
	// 处理异常响应
	if strings.Contains(response.Request.URL, "/api/") {
		// 检查 Content-Type
		contentType := response.Header().Get("Content-Type")
		if !strings.Contains(contentType, "application/json") {
			return uniapi.RaiseForStatus(response)
		}

		// 解析通用响应体
		var resp BestdoriAPIResponse
		if err := json.Unmarshal(response.Body(), &resp); err != nil {
			return err
		}
		if resp.Result == nil {
			if response.IsError() {
				return &uniapi.ResponseStatusError{Response: response}
			}
			return nil
		} else if !*resp.Result {
			code := resp.Code
			return &RequestFailedError{
				ResponseError: &uniapi.ResponseError{Response: response},
				Code:          code,
			}
		}
		return uniapi.RaiseForStatus(response)
	} else if strings.Contains(response.Request.URL, "/assets/") || strings.Contains(response.Request.URL, "/res/") {
		// 检查 Content-Type
		if strings.Contains(response.Header().Get("Content-Type"), "text/html") {
			return &AssetsNotExistError{Response: response}
		}
		return uniapi.RaiseForStatus(response)
	}
	return uniapi.RaiseForStatus(response)
}
