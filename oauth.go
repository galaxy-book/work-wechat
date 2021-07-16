package work

import (
	"encoding/json"
	"fmt"
)

func GetUserInfo3RD(suiteAccessToken, code string) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/service/getuserinfo3rd?suite_access_token=%s&code=%s", suiteAccessToken, code)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpGet),
	)
}

func GetUserDetail3RD(suiteAccessToken, userTicket string) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/service/getuserdetail3rd?suite_access_token=%s", suiteAccessToken)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpPost),
		WitchBody(func() (bytes []byte, e error) {
			req := map[string]interface{}{
				"user_ticket": userTicket,
			}
			jsonInfo, err := json.Marshal(req)
			if err != nil {
				return nil, err
			}
			return jsonInfo, nil
		}),
	)
}

func GetLoginInfo(providerAccessToken, authCode string) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/service/get_login_info?access_token=%s", providerAccessToken)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpPost),
		WitchBody(func() (bytes []byte, e error) {
			req := map[string]interface{}{
				"auth_code": authCode,
			}
			jsonInfo, err := json.Marshal(req)
			if err != nil {
				return nil, err
			}
			return jsonInfo, nil
		}),
	)
}
