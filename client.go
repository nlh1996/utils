package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

var client *http.Client

func init() {
	// 不验证证书合法
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}
}

// HttpPost .
func HttpPost(url string, data interface{}) (string, error) {
	bytesData, err := json.Marshal(data)
	if err != nil {
		return "解析出错", err
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	resp, err := client.Do(request)
	if err != nil {
		return "请求失败", err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "读取失败", err
	}
	if resp.StatusCode != 200 {
		err := errors.New(string(respBytes))
		return "请求失败", err
	}
	return string(respBytes), nil
}

// HttpGet .
func HttpGet(url string) []byte {
	resp, err := client.Get(url)
	if err != nil {
		log.Println(err)
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil
	}

	if resp != nil {
		resp.Body.Close()
	}
	return body
}
