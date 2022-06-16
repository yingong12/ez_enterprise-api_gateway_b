package controller

import (
	"api_gateway_b/http/buz_code"
	"api_gateway_b/http/middleware"
	"api_gateway_b/library/env"
	"api_gateway_b/providers"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ForwardCompanyService(ctx *gin.Context) {
	//企业更新。 转发
	if ctx.Request.URL.Path == "/enterprise/update" || ctx.Request.URL.Path == "enterprise/update" {
		STDWrapperRaw(EnterpriseUpdate)(ctx)
		return
	}
	Proxy(ctx, env.GetStringVal("LB_COMPANY_SERVICE"))
}

//更新企业
func EnterpriseUpdate(ctx *gin.Context) (res RawResponse, err error) {
	clientBody := &bytes.Reader{}
	v, _ := ctx.Get("auth_info")
	appID := v.(*middleware.AuthInfo).AppID
	if appID == "" {
		ctx.JSON(200, gin.H{
			"code": buz_code.CODE_AUTH_FAILED,
			"msg":  "该用户还没绑定企业或机构",
		})
		return
	}
	if err = func() error {
		body := ctx.Request.Body
		bodyMap := map[string]interface{}{}
		err := json.NewDecoder(body).Decode(&bodyMap)
		if err != nil {
			return fmt.Errorf("decode failed")
		}
		j, _ := json.Marshal(bodyMap)
		clientBody = bytes.NewReader(j)
		return nil
	}(); err != nil {
		//直接把错误抛给后端
		res = []byte(err.Error())
		return
	}
	//发送请求
	client := providers.HttpClientCompanyService
	//从url里拿appID
	URL := fmt.Sprintf("%s/enterprise/%s", client.BaseURL, appID)
	request, err := http.NewRequest("PUT", URL, clientBody)
	if err != nil {
		return
	}
	rsp, err := client.Do(request)
	if err != nil {
		return
	}
	//解析rsp
	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return
	}
	//
	res = data
	return
}
