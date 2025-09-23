package dto

// 空结构体，仅表示字段名存在
type EmptyStruct map[string]struct{}

// 服务器 ID
type Server int

const (
	ServerJP Server = 1
	ServerEN Server = 2
	ServerTW Server = 3
	ServerCN Server = 4
	ServerKR Server = 5
)

// 服务器名称
type ServerName string

const (
	ServerNameJP ServerName = "jp"
	ServerNameEN ServerName = "en"
	ServerNameTW ServerName = "tw"
	ServerNameCN ServerName = "cn"
	ServerNameKR ServerName = "kr"
)
