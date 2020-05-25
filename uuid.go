package utils

import (
	"strconv"
	"sync/atomic"
	"time"
)

var index int64

// GetUID 获取唯一uid
func GetUID(tag string) string {
	duration, _ := time.ParseDuration("8h")
	now := time.Now().UTC().Add(duration)
	strTime := now.Format("060102030405")
	// if index >= 99 {
	// 	index = 0
	// }
	atomic.AddInt64(&index, 1)
	uid := tag + strTime + strconv.FormatInt(index, 10)
	return uid
}
