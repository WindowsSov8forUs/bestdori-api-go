package dto

type ComicInfo struct {
	AssetBundleName string       `json:"assetBundleName"` // 漫画资源库名
	Title           []*string    `json:"title"`           // 漫画标题 定长列表
	SubTitle        []*string    `json:"subTitle"`        // 漫画副标题 定长列表
	PublicStartAt   []*Timestamp `json:"publicStartAt"`   // 漫画公开时间时间戳 定长列表。若漫画为服务器开放时即开放则此字段值为 1 ，否则为时间戳字符串
	CharacterId     []int        `json:"characterId"`     // 漫画相关角色 ID 列表
}

type ComicsAll5 map[string]ComicInfo
