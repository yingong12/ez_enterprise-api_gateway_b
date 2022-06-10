package middleware

import (
	"api_gateway_b/http/buz_code"
	"api_gateway_b/providers"
	"bytes"
	"encoding/json"
	"errors"
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
		token := c.Request.Header.Get("b_access_token")
		code, msg, appID, uID, err := openAuth(token)
		log.Println("进入鉴权中间件")
		if err != nil {
			//网络错误
			c.JSON(http.StatusOK, gin.H{
				"code": buz_code.CODE_SERVER_ERROR,
				"msg":  "服务器内部错误",
			})
			c.Abort()
			return
		}
		//鉴权失败
		if code != int(buz_code.CODE_OK) {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  msg,
			})
			c.Abort()
		}
		log.Println("鉴权中间件解析")
		info := &AuthInfo{
			AppID: appID,
			UID:   uID,
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

type BaseRsp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
type RspAuthInfo struct {
	BaseRsp
	Data struct {
		UID       string `json:"uid"`       //b端用户id
		AppID     string `json:"app_id"`    //appID
		ExpiresAt string `json:"expire_at"` //过期时间

	} `json:"data"`
}

func openAuth(token string) (code int, msg, appID, uID string, err error) {
	//TODO mock
	client := providers.HttpClientAccountService
	req, err := http.NewRequest("GET", client.BaseURL+"/auth/check", nil)
	req.Header.Set("b_access_token", token)
	rsp, err := client.Do(req)
	if err != nil {
		return
	}
	if rsp.StatusCode != http.StatusOK {
		err = errors.New("http_code not 200")
		return
	}
	res := RspAuthInfo{}
	err = json.NewDecoder(rsp.Body).Decode(&res)
	if err != nil {
		return
	}
	code = res.Code
	msg = res.Msg
	uID = res.Data.UID
	appID = res.Data.AppID
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
