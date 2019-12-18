package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// 发送GET请求
// url:请求地址
// params:请求参数
// headers:请求头信息
// content:请求返回的内容
// 默认5s超时

func HttpGet(url string, params, headers map[string]string) (content string) {
	req, _ := http.NewRequest("GET", url, nil)

	for k, v := range headers {
		req.Header.Add(k, v)
	}
	q := req.URL.Query()
	for key, val := range params {
		q.Add(key, val)
	}
	req.URL.RawQuery = q.Encode()
	defer req.Body.Close()
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	content = string(result)
	log.Println("=============发送的请求:============")
	log.Println(req.URL.String())
	log.Println("=============返回的结果:============")
	log.Println(content)
	return
}

func HttpPost(rawUrl string, params interface{}, headers map[string]string) (respBytes []byte, err error) {
	bytesData, err := json.Marshal(params)
	if err != nil {
		log.Println(err.Error())
		return
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", rawUrl, reader)
	if err != nil {
		log.Println(err.Error())
		return
	}
	//request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	if _, ok := headers["Content-Type"]; !ok {
		headers["Content-Type"] = "application/json;charset=UTF-8"
	}
	for k, v := range headers {
		request.Header.Add(k, v)
	}
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
		return
	}
	respBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("=============发送的请求:============")
	log.Println(rawUrl)
	log.Println(string(bytesData))
	log.Println("=============返回的结果:============")
	log.Println(string(respBytes))
	return respBytes, nil
}
