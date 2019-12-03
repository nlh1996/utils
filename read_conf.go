package utils

import (
	"encoding/json"
	"io/ioutil"
)

/*
	读取全局配置
*/

// ReadFile .
func ReadFile(data interface{}) {
	filedata, err := ioutil.ReadFile("conf/conf.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(filedata, data); err != nil {
		panic(err)
	}
}
