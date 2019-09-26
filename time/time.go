package time

import (
	"time"
)

// GetMillisecond 获取毫秒数
func GetMillisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
