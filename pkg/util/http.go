package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
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

func HttpPost(url string, param interface{}, header map[string]string) (*HttpResp, error) {
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
	req.Header.Set("Content-Type", "application/json")

	for k, v := range header {
		req.Header.Set(k, v)
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
