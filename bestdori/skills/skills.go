package skills

import (
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/dto"
	"github.com/WindowsSov8forUs/bestdori-api-go/bestdori/endpoints"
	"github.com/WindowsSov8forUs/bestdori-api-go/uniapi"
)

// GetAll2 获取总技能简易描述信息
func GetAll2(api *uniapi.UniAPI) (*dto.SkillsAll2, error) {
	endpoint := endpoints.SkillsAll(2)
	return uniapi.Get[dto.SkillsAll2](api, endpoint, nil)
}

// GetAll5 获取总技能简洁信息
func GetAll5(api *uniapi.UniAPI) (*dto.SkillsAll5, error) {
	endpoint := endpoints.SkillsAll(5)
	return uniapi.Get[dto.SkillsAll5](api, endpoint, nil)
}

// GetAll10 获取总技能详细信息
func GetAll10(api *uniapi.UniAPI) (*dto.SkillsAll10, error) {
	endpoint := endpoints.SkillsAll(10)
	return uniapi.Get[dto.SkillsAll10](api, endpoint, nil)
}
