package utils

import (
	"log"
	"strconv"
)

// StringToInt .
func StringToInt(data string) int {
	tag, err := strconv.Atoi(data)
	if err != nil {
		log.Println(err)
		return 0
	}
	return tag
}

// StringToInt64 .
func StringToInt64(data string) int64 {
	tag, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		log.Println(err)
		return 0
	}
	return tag
} 

// IntToString .
func IntToString(data int) string {
	return strconv.Itoa(data)
}

// Int64ToString .
func Int64ToString(data int64) string {
	return strconv.FormatInt(data,10)
}
