package library

import (
	"io"
	"net"
	"net/http"
	"time"
)

type HttpClient struct {
	*http.Client
	BaseURL string
}
type HttpClientConfig struct {
	Name                      string //名称
	DialTimeoutSecond         int    // 连接超时
	DialKeepAliveSecond       int    //开启长连接
	MaxIdleConnections        int    //最大空闲连接数
	MaxIdleConnectionsPerHost int    //单Host最大空闲连接数
	IdleConnTimeoutSecond     int    // 空闲连接超时
	Receiver                  **HttpClient
	BaseURL                   string
}

func NewHttpClient(config *HttpClientConfig) (httpClient *HttpClient) {
	httpClient = &HttpClient{
		Client: &http.Client{
			Transport: &http.Transport{
				DisableKeepAlives:   config.DialKeepAliveSecond < 0,
				MaxIdleConns:        config.MaxIdleConnections,
				MaxIdleConnsPerHost: config.MaxIdleConnectionsPerHost,
				IdleConnTimeout:     time.Duration(config.IdleConnTimeoutSecond) * time.Second,
				Proxy:               http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   time.Duration(config.DialTimeoutSecond) * time.Second,
					KeepAlive: time.Duration(config.DialKeepAliveSecond) * time.Second,
				}).DialContext,
			},
		},
	}
	return
}

func (c *HttpClient) PostJson(url string, body io.Reader) (resp *http.Response, err error) {
	return c.Post(url, "application/json", body)
}
