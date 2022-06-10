package controller

import (
	"api_gateway_b/http/middleware"
	"api_gateway_b/providers"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ForwardCompanyService(ctx *gin.Context) {
	ctx.Writer.Write([]byte("redirecting to company_service"))
}

//更新企业
func UpdateEnterprise(ctx *gin.Context) (response RawResponse, err error) {
	//解析并重新封装请求体
	clientBody := &bytes.Reader{}
	authInfo, _ := middleware.GetAuthInfo(ctx)
	if err = func() error {
		body := ctx.Request.Body
		bodyMap := map[string]interface{}{}
		err := json.NewDecoder(body).Decode(&bodyMap)
		if err != nil {
			return fmt.Errorf("decode failed")
		}
		bodyMap["uid"] = authInfo.UID
		j, _ := json.Marshal(bodyMap)
		clientBody = bytes.NewReader(j)
		return nil
	}(); err != nil {
		return
	}
	//发送请求
	client := providers.HttpClientCompanyService
	appID := authInfo.AppID
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
	response = data
	return
}
