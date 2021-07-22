package work

type GetDeptListResp struct {
	ErrCode    int        `json:"errcode"`
	ErrMsg     string     `json:"errmsg"`
	Department []DeptInfo `json:"department"`
}

type DeptInfo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	NameEn   string `json:"name_en"`
	ParentId int    `json:"parentid"`
	Order    int    `json:"order"`
}

type GetDeptMemberListResp struct {
	ErrCode  int               `json:"errcode"`
	ErrMsg   string            `json:"errmsg"`
	UserList []UserInfoDetails `json:"userList"`
}
