package time

import (
	"time"
)

// Time 格式化当前时间
func Time() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// TimTransfer 转换时间戳
func TimTransfer(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}
