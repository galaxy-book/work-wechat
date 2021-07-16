package work

type GetUserInfo3rdResp struct {
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
	CorpId     string `json:"CorpId"`
	UserId     string `json:"UserId"`
	DeviceId   string `json:"DeviceId"`
	UserTicket string `json:"user_ticket"`
	ExpiresIn  int64  `json:"expires_in"`
	OpenUserId string `json:"open_userid"`
}

type GetUserDetail3rdResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	CorpId  string `json:"CorpId"`
	UserId  string `json:"UserId"`
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Avatar  string `json:"avatar"`
}

type GetLoginInfoResp struct {
	ErrCode  int              `json:"errcode"`
	ErrMsg   string           `json:"errmsg"`
	UserType int              `json:"usertype"`
	UserInfo UserInfo         `json:"user_info"`
	CorpInfo CorpInfo         `json:"corp_info"`
	Agent    []LoginAgentInfo `json:"agent"`
	AuthInfo LoginAuthInfo    `json:"authInfo"`
}

type CorpInfo struct {
	CorpId string `json:"corpid"`
}

type LoginAgentInfo struct {
	AgentId  int `json:"agentid"`
	AuthType int `json:"auth_type"`
}

type LoginAuthInfo struct {
	Department LoginAuthInfoDepartment `json:"department"`
}

type LoginAuthInfoDepartment struct {
	Id       int  `json:"id"`
	Writable bool `json:"writable"`
}
