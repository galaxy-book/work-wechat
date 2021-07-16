package work

type GetCorpAuthInfoResp struct {
	respCommon
	GetCorpAuthInfo
}

type GetCorpAuthInfo struct {
	AuthCorpInfo   AuthCorpInfo   `json:"auth_corp_info"`
	AuthInfo       AuthInfo       `json:"auth_info"`
	DealerCorpInfo DealerCorpInfo `json:"dealer_corp_info"`
}
