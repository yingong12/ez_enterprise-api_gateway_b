package controller

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

//直接转发
func Proxy(ctx *gin.Context, host string) {
	u := &url.URL{}
	//转发的协议，如果是https，写https.
	u.Scheme = "http"
	u.Host = host
	proxy := httputil.NewSingleHostReverseProxy(u)

	//重写出错回调
	proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
		log.Printf("http: proxy error: %v", err)
		ret := fmt.Sprintf("http proxy error %v", err)
		//写到body里
		rw.Write([]byte(ret))
	}
	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}
