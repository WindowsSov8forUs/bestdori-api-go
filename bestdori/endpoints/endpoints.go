package endpoints

import (
	"strings"
	"sync"
)

const (
	json = ".json"
	png  = ".png"
	mp3  = ".mp3"
	svg  = ".svg"
)

var builderPool = sync.Pool{
	New: func() any {
		builder := &strings.Builder{}
		builder.Grow(256) // 预分配 256 字节容量
		return builder
	},
}

func getBuilder() *strings.Builder {
	return builderPool.Get().(*strings.Builder)
}

func putBuilder(builder *strings.Builder) {
	builder.Reset()
	builderPool.Put(builder)
}
