package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// ReadFile .
// 读取json配置文件
// data 结构体指针
// path 配置文件相对路径 例如"conf/conf.json"
func ReadFile(data interface{}, path string) error {
	if confFileExists, err := PathExists(path); confFileExists != true {
		return err
	}

	filedata, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(filedata, data)
}

// PathExists 判断一个文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}
