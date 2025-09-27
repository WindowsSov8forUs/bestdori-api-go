package bestdori

import (
	"fmt"
	"strings"

	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

type NotExistError struct {
	uniapi.ResponseError
	Target string
}

func (e *NotExistError) Error() string {
	return fmt.Sprintf("%s does not exist", e.Target)
}

type ServerNotAvailableError struct {
	uniapi.ResponseError
	Target string
	Server dto.ServerName
}

func (e *ServerNotAvailableError) Error() string {
	return fmt.Sprintf("%s is not available on server %s", e.Target, e.Server)
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
		return 0, fmt.Errorf("unknown server name: %s", serverName)
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
		return "", fmt.Errorf("unknown server id: %d", serverId)
	}
}

func RemoveURLPrefix(url string) string {
	if strings.HasPrefix(url, "https://bestdori.com") {
		return strings.TrimPrefix(url, "https://bestdori.com")
	}
	return url
}
