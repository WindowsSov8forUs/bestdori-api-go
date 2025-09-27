package dto

import (
	"encoding/json"
	"strconv"
	"time"
)

// 空结构体，仅表示字段名存在
type EmptyStruct map[string]struct{}

// 服务器 ID
type Server int

const (
	ServerJP Server = iota
	ServerEN
	ServerTW
	ServerCN
	ServerKR
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

// 时间戳类型，用于处理 Bestdori 的奇异时间戳处理
type Timestamp struct {
	IsBuiltIn bool      // 是否为该服务器开放时即存在的内容
	Time      time.Time // 公开时间
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	if str == "1" {
		t.IsBuiltIn = true
		t.Time = time.Time{}
	} else {
		timestamp, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return err
		}

		t.IsBuiltIn = false
		t.Time = time.UnixMilli(timestamp)
	}
	return nil
}
