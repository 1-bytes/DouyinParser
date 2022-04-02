package fetcher

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// JsonRequest 以 Json 格式获取网页内容.
func JsonRequest(method string, api string, body string, header *http.Header) (*http.Response, error) {
	//u, err := url.Parse(api)
	//if err != nil {
	//	return nil, err
	//}
	// 加入对应的 header
	//if header.Get("Host") != u.Host {
	//	header.Add("Host", u.Host)
	//	header.Add("Content-Length", strconv.Itoa(len(body)))
	//	header.Add("Content-type", "application/json")
	//}
	// 构造请求
	req, err := http.NewRequest(method, api, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := sendRequest(req, header)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed, unintended status code: %d", resp.StatusCode)
	}
	return resp, nil
}

// sendRequest 发出请求.
func sendRequest(req *http.Request, header *http.Header) (*http.Response, error) {
	if header.Get("User-Agent") == "" {
		header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) "+
			"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36")
	}
	req.Header = *header
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	return client.Do(req)
}
