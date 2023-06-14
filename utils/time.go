package utils

import (
	"time"
)

// Time 格式化当前时间
func Time() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
