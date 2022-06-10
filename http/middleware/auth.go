package middleware

import (
	"api_gateway_b/http/buz_code"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//登录态信息, 请求context公用
type AuthInfo struct {
	AppID string `json:"app"`
	UID   string `json:"uid"`
}

//从请求体内拿登录态信息
func GetAuthInfo(ctx *gin.Context) (info *AuthInfo, ok bool) {
	d, ok := ctx.Get("auth_info")
	info = d.(*AuthInfo)
	return
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//
		code, msg, appID, uID, err := openAuth()
		log.Println("进入鉴权中间件")
		if err != nil {
			//网络错误
			c.JSON(http.StatusOK, gin.H{
				"code": buz_code.CODE_SERVER_ERROR,
				"msg":  "服务器内部错误",
			})
			c.Abort()
		}
		//鉴权失败
		if code != buz_code.CODE_OK {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  msg,
			})
			c.Abort()
		}
		log.Println("鉴权中间件解析", appID, uID)
		info := &AuthInfo{
			AppID: appID,
		}
		c.Set("auth_info", info)
		c.Next()

	}
}

//HeaderInjector injects the header to the request
func HeaderInjector() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}

func openAuth() (code buz_code.Code, msg, appID, uID string, err error) {
	//TODO mock
	appID = "app1"
	return
}

func RequestLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//后置
		before := time.Now()
		bodyStream := copyBody(ctx)
		ctx.Next()

		after := time.Now()
		//配置到单独logger
		log.Printf("[response_time]:%dms, [end_point]:%s, [method]:%s, [body]:%s\n", after.Sub(before).Milliseconds(), ctx.Request.URL.Path, ctx.Request.Method, removeJSONIndent(bodyStream))
	}
}
func removeJSONIndent(input []byte) (output []byte) {
	bstd := map[string]interface{}{}
	json.Unmarshal(input, &bstd)
	output, _ = json.Marshal(bstd)
	return
}

//copy request body
func copyBody(ctx *gin.Context) (buf []byte) {
	dist := &bytes.Buffer{}
	//从src读写到dst
	trdr := io.TeeReader(ctx.Request.Body, dist)
	//注意 这里err了请求就丢了
	buf, _ = ioutil.ReadAll(trdr)
	ctx.Request.Body = io.NopCloser(dist)
	return
}
