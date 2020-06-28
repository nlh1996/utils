package utils

import (
	"context"
	"strconv"
	"time"
)

// GetTimeStr 将时间格式化为字符串（东八区）.
func GetTimeStr() string {
	//世界时间转北京时间
	duration, _ := time.ParseDuration("8h")
	now := time.Now().UTC().Add(duration)
	strTime := now.Format("2006-01-02 15:04:05")
	return strTime
}

// GetDateStr 获取当天日期（东八区）
func GetDateStr() string {
	duration, _ := time.ParseDuration("8h")
	now := time.Now().UTC().Add(duration)
	strTime := now.Format("2006-01-02")
	return strTime
}

// ComputeAge 根据身份证计算年龄
func ComputeAge(str string) int {
	if len(str) != 18 {
		return 0
	}
	year, _ := strconv.Atoi(str[6:10])
	mounth, _ := strconv.Atoi(str[10:12])
	day, _ := strconv.Atoi(str[12:14])
	date := GetDateStr()
	nowYear, _ := strconv.Atoi(date[0:4])
	nowMounth, _ := strconv.Atoi(date[5:7])
	nowDay, _ := strconv.Atoi(date[8:10])
	if year >= nowYear {
		return 0
	}
	age := nowYear - year
	if nowMounth < mounth {
		age--
	}
	if nowMounth == mounth && nowDay < day {
		age--
	}
	return age
}

// ComputeSecond 计算时间戳与当前时间戳间隔秒数
func ComputeSecond(d int64) int64 {
	now := time.Now().Unix()
	return now - d
}

// Timer 定时器 callback回调函数 d触发间隔 count循环次数 -1无限循环
func Timer(callback func(), d time.Duration, count int, ctx context.Context) {
	if count == -1 {
		for range time.Tick(d) {
			select {
			case <-ctx.Done():
				return
			default:
				callback()
			}
		}
	} else {
		for i := 0; i < count; i++ {
			time.Sleep(d)
			callback()
		}
	}
}
