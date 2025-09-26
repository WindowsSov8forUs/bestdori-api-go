package bestdori

import (
	"fmt"

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
