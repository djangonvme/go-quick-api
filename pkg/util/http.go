package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type HttpResp struct {
	// *http.Response
	StatusCode int
	Body       []byte
}

func (h *HttpResp) StatusOk() bool {
	return h.StatusCode == http.StatusOK
}
func (h *HttpResp) Decode(dst interface{}) error {
	if len(h.Body) == 0 {
		return nil
	}
	return json.Unmarshal(h.Body, &dst)
}

func HttpPost(url string, headers map[string]string, param interface{}) (*HttpResp, error) {
	url = strings.ReplaceAll(url, " ", "")
	var postBytes []byte
	rawText, ok := param.(string)
	if ok {
		postBytes = []byte(rawText)
	} else {
		bte, err := json.Marshal(param)
		if err != nil {
			return nil, err
		}
		postBytes = bte
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postBytes))
	if err != nil {
		return nil, err
	}
	if headers != nil && len(headers) > 0 {
		if _, ok = headers["Content-Type"]; !ok {
			req.Header.Set("Content-Type", "application/json")
		}
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &HttpResp{Body: body, StatusCode: resp.StatusCode}, nil
}

func HttpGet(reqUrl string, params map[string]string, headers map[string]string) (*HttpResp, error) {
	if params != nil && len(params) > 0 {
		arg := url.Values{}
		for k, v := range params {
			arg.Add(k, v)
		}
		reqUrl += "?" + arg.Encode()
	}
	req, _ := http.NewRequest("GET", reqUrl, nil)
	if headers != nil && len(headers) > 0 {
		if _, ok := headers["Content-Type"]; !ok {
			req.Header.Set("Content-Type", "application/json")
		}
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &HttpResp{Body: body, StatusCode: resp.StatusCode}, nil
}

func RetryWithTimeoutDo(timeout time.Duration, fn func() error) {
	err := fn()
	if err == nil {
		return
	}
	elapse := 1 * time.Second
	count := 1
	for {
		count++
		select {
		case <-time.After(timeout):
			return
		case <-time.Tick(elapse):
			if err = fn(); err != nil {
				elapse += time.Duration(count) * time.Second
			} else {
				return
			}
		}
	}
}
