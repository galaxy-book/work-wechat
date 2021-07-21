package work

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

func GetProviderAccessTokenAction(providerCorpId string, providerSecret string) Action {
	reqUrl := BaseWeWorkUrl + "/cgi-bin/service/get_provider_token"
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpPost),
		WitchBody(func() (bytes []byte, e error) {
			var req = reqGetProviderToken{
				CorpId:         providerCorpId,
				ProviderSecret: providerSecret,
			}
			jsonInfo, err := json.Marshal(req)
			if err != nil {
				return nil, err
			}
			return jsonInfo, nil
		}),
	)
}

func GetSuitAccessTokenAction(suiteId string, suiteSecret string, suiteTicket string) Action {
	reqUrl := BaseWeWorkUrl + "/cgi-bin/service/get_suite_token"
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpPost),
		WitchBody(func() (bytes []byte, e error) {
			var req = reqGetSuiteToken{
				SuiteID:     suiteId,
				SuitSecret:  suiteSecret,
				SuiteTicket: suiteTicket,
			}
			jsonInfo, err := json.Marshal(req)
			if err != nil {
				return nil, err
			}
			return jsonInfo, nil
		}),
	)
}

func GetCorpAccessTokenAction(suitAccessToken string, corpId string, permanentCode string) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/service/get_corp_token?suite_access_token=%s", suitAccessToken)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpPost),
		WitchBody(func() (bytes []byte, e error) {
			var req = reqGetCorpToken{
				AuthCorpID:    corpId,
				PermanentCode: permanentCode,
			}
			jsonInfo, err := json.Marshal(req)
			if err != nil {
				return nil, err
			}
			return jsonInfo, nil
		}),
	)
}

func GetCorpAuthInfoAction(suitAccessToken string, corpId string, permanentCode string) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/service/get_auth_info?suite_access_token=%s", suitAccessToken)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpPost),
		WitchBody(func() (bytes []byte, e error) {
			var req = reqGetCorpToken{
				AuthCorpID:    corpId,
				PermanentCode: permanentCode,
			}
			jsonInfo, err := json.Marshal(req)
			if err != nil {
				return nil, err
			}
			return jsonInfo, nil
		}),
	)
}

/**
 * @Description:
 * @author:21
 * @receiver w
 * @return *RespGetProviderToken
 * @return error
 */
func (w *WorkWechat) GetProviderAccessToken() (*RespGetProviderToken, error) {
	if len(w.ProviderCorpID) < 1 {
		return nil, errors.New("设置ProviderCorpID出错")
	}

	if len(w.ProviderSecret) < 1 {
		return nil, errors.New("设置ProviderSecret出错")
	}

	var resp = &RespGetProviderToken{}
	err := w.Scan(context.Background(),
		GetProviderAccessTokenAction(w.ProviderCorpID, w.ProviderSecret),
		resp,
	)

	if err != nil {
		return nil, err
	}

	if resp.respCommon.ErrCode != 0 {
		return nil, errors.New(resp.respCommon.ErrMsg)
	}
	return resp, nil
}

/**
 * @Description:获取第三方应用access_token
 * @author:ljj
 * @receiver w
 * @return *RespGetSuiteToken
 * @return error
 */
func (w *WorkWechat) GetSuiteAccessToken() (*RespGetSuiteToken, error) {
	if len(w.SuiteID) < 1 {
		return nil, errors.New("设置SuiteID出错")
	}

	if len(w.SuiteSecret) < 1 {
		return nil, errors.New("设置SuiteSecret出错")
	}

	var resp = &RespGetSuiteToken{}
	err := w.Scan(context.Background(),
		GetSuitAccessTokenAction(w.SuiteID, w.SuiteSecret, w.SuiteTicket),
		resp,
	)
	if err != nil {
		return nil, err
	}
	if resp.respCommon.ErrCode != 0 {
		return nil, errors.New(resp.respCommon.ErrMsg)
	}
	return resp, nil
}

/**
 * @Description:获取授权企业应用access_token
 * @author:ljj
 * @receiver w
 * @return *RespGetCorpToken
 * @return error
 */
func (w *WorkWechat) GetCorpAccessToken() (*RespGetCorpToken, error) {
	if len(w.CorpId) < 1 {
		return nil, errors.New("设置CorpId出错")
	}
	if len(w.PermanentCode) < 1 {
		return nil, errors.New("设置PermanentCode出错")
	}
	suiteAccessTokenResp, err := w.GetSuiteAccessToken()
	if err != nil {
		return nil, err
	}
	var resp = &RespGetCorpToken{}
	err = w.Scan(context.Background(),
		GetCorpAccessTokenAction(suiteAccessTokenResp.SuiteAccessToken, w.CorpId, w.PermanentCode),
		resp,
	)
	if err != nil {
		return nil, err
	}
	if resp.respCommon.ErrCode != 0 {
		return nil, errors.New(resp.respCommon.ErrMsg)
	}
	return resp, nil
}
