package work

import (
	"encoding/json"
	"fmt"
)

func SendMsg(accessToken string, req SendMsgReq) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/message/send?access_token=%s", accessToken)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpPost),
		WitchBody(func() (bytes []byte, e error) {
			return json.Marshal(req)
		}),
	)
}
