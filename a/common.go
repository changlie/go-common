package a

import (
	"github.com/google/uuid"
	"strings"
)

func Uuid() string {
	return strings.ReplaceAll(UuidV4(), "-", "")
}

func UuidV4() string {
	// V4 基于随机数
	u4 := uuid.New()
	return u4.String()
}
