package dto

type CostumesAll5Info struct {
	CharacterId     int       `json:"characterId"`     // 服装对应角色 ID
	AssetBundleName string    `json:"assetBundleName"` // 服装资源库名
	Description     []*string `json:"description"`     // 服装描述 定长列表
	PublishedAt     []*string `json:"publishedAt"`     // 服装上线时间时间戳 定长列表
}

type CostumesAll5 map[string]CostumesAll5Info

type CostumeInfo struct {
	CostumesAll5Info `json:",inline"`
	SdResourceName   string    `json:"sdResourceName"` // 服装 LIVESD 资源库名
	HowToGet         []*string `json:"howToGet"`       // 服装获取方法 定长列表
	Cards            []int     `json:"cards"`          // 服装对应卡牌 ID 列表
}
