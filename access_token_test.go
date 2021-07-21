package work

import (
	"context"
	"testing"
)

func TestAccessToken_GetCorpAccessToken(t *testing.T) {
	wechatSDK := NewWorkWechat(Config{
		SuiteID:     "xx",
		SuiteSecret: "xx",
		SuiteTicket: "xx",
	})
	suiteAccessToken, err := wechatSDK.GetSuiteAccessToken()
	if err != nil {
		t.Fatal(err)
	}

	action := NewGetPermanentCode(suiteAccessToken.SuiteAccessToken, "xxx")
	body, err := action.DoRequest(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(body)
}
