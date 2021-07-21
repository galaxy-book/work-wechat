package work

import (
	"context"
	"testing"
)

func TestAccessToken_GetCorpAccessToken(t *testing.T) {
	wechatSDK := NewWorkWechat(Config{
		SuiteID:     "ww9b85ae8ff033ee89",
		SuiteSecret: "BCLomiIeq8je52OqsXusskBMSMO8LSLnuIxpxMnfhrc",
		SuiteTicket: "zHIaXmHYu-UWu_hOXICtNyp5OggHhA6sthSRWAoUfRe4t2OeZiLHoRG1-_3LWVON",
	})
	suiteAccessToken, err := wechatSDK.GetSuiteAccessToken()
	if err != nil {
		t.Fatal(err)
	}

	action := NewGetPermanentCode(suiteAccessToken.SuiteAccessToken, "ikr5f_hkVQx5GMBLI0FZl60ZulcSWCi0Mc9oULynqu2YIG8ABuywVvhl5LYY663XwHcvzbNgrwyFpG_e2VnzMkQzuKNcjPj0Zp1cHBlM5XQ")
	body, err := action.DoRequest(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(body)
}
