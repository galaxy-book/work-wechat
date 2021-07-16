package work

import (
	"fmt"
)

func GetDeptList(accessToken, id string) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/department/list?access_token=%s&id=%s", accessToken, id)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpGet),
	)
}

func GetDeptMemberList(accessToken, departmentId string, fetchChild int) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/user/list?access_token=%s&department_id=%s&fetch_child=%d", accessToken, departmentId, fetchChild)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpGet),
	)
}
