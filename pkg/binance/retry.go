package binance

import (
	"io"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

// 全局复用的重试客户端
var client *retryablehttp.Client

func init() {
	client = retryablehttp.NewClient()
	client.RetryMax = 5
	client.HTTPClient.Timeout = 5 * time.Second
	client.Logger = nil
}

// DoGet 是通用 GET 请求封装，返回响应体内容
func DoGet(url string, headers map[string]string) ([]byte, error) {
	req, err := retryablehttp.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
