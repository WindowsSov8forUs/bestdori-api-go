package dto

// 乐队信息
type BandsInfo struct {
	BandName []*string `json:"bandName"` // 乐队名称 定长列表
}

// 总乐队信息
type BandsAll1 map[string]BandsInfo

// 主要乐队信息
type BandsMain1 BandsAll1
