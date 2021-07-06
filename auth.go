package work

import (
	"errors"
	"github.com/gogf/gf/util/gconv"
)

type auth struct {
	workWechat workWechat
}

func (w workWechat) NewAuth() *auth {
	return &auth{
		workWechat: w,
	}
}

// GetPreAuthCode 获取预授权码 https://work.weixin.qq.com/api/doc/90001/90143/90601
func (a *auth) GetPreAuthCode () (respGetPreAuthCode, error) {

	suitAccessToken := a.workWechat.NewAccessToken().getSuitAccessToken()

	var req = reqGetPreAuthCode{
		SuiteAccessToken: suitAccessToken,
	}
	var respGetPreAuthCode = respGetPreAuthCode{}

	resp, err := HttpClient.httpGet("/cgi-bin/service/get_pre_auth_code", req)
	if err != nil {
		return respGetPreAuthCode, err
	}
	if resp == "" {
		return respGetPreAuthCode, errors.New("请求错误")
	}
	err = gconv.Struct(resp, respGetPreAuthCode)
	if err != nil {
		return respGetPreAuthCode, err
	}

	return respGetPreAuthCode, nil
}

// GetPermanentCode 获取企业永久授权码 https://work.weixin.qq.com/api/doc/90001/90143/90603
func (a *auth) GetPermanentCode (authCode string) (respGetPermanentCode, error) {

	suitAccessToken := a.workWechat.NewAccessToken().getSuitAccessToken()

	var req = reqGetPermanentCode{
		AuthCode: authCode,
	}
	var respGetPermanentCode = respGetPermanentCode{}

	resp, err := HttpClient.httpPost("/cgi-bin/service/get_permanent_code?" + suitAccessToken, req)
	if err != nil {
		return respGetPermanentCode, err
	}
	if resp == "" {
		return respGetPermanentCode, errors.New("请求错误")
	}
	err = gconv.Struct(resp, respGetPermanentCode)
	if err != nil {
		return respGetPermanentCode, err
	}

	return respGetPermanentCode, nil
}
