package work

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

// 外部联系人相关
type externalContact struct {
	workWechat WorkWechat
}

func (w WorkWechat) NewExternalContact() *externalContact {
	return &externalContact{
		w,
	}
}

// 获取客户列表
func NewGetExternalContactList(accessToken string, userId string) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/externalcontact/list?access_token=%s&userid=%s", accessToken, userId)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpGet),
	)
}

/**
 * @Description: 获取客户列表
 * @author:ljj
 * @receiver w
 * @param userId string 企业成员的userid
 * @return *ExternalContactList
 * @return error
 */
func (e *externalContact) GetExternalContactList(userId string) (*ExternalContactList, error) {

	corpAccessTokenResp, err := e.workWechat.GetCorpAccessToken()
	if err != nil {
		return nil, err
	}
	corpAccessToken := corpAccessTokenResp.AccessToken

	opt := &ExternalContactList{}
	err = e.workWechat.Scan(context.Background(), NewGetExternalContactList(corpAccessToken, userId), opt)
	if err != nil {
		return nil, err
	}
	if opt.ErrCode != 0 {
		return nil, errors.New("设置授权配置失败")
	}
	return opt, nil
}

// 获取客户详细信息
func NewGetExternalContactUserInfo(accessToken string, externalUserId string, cursors string) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/externalcontact/get?access_token=%s&external_userid=%s&cursor=%s", accessToken, externalUserId, cursors)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpGet),
	)
}

/**
 * @Description: 获取客户详细信息
 * @author:ljj
 * @receiver w
 * @param externalUserId string 外部联系人的userid，注意不是企业成员的帐号
 * @param cursor string 上次请求返回的next_cursor
 * @return *ExternalContactUserInfo
 * @return error
 */
func (e *externalContact) GetExternalContactUserInfo(externalUserId string, cursor string) (*ExternalContactUserInfo, error) {

	corpAccessTokenResp, err := e.workWechat.GetCorpAccessToken()
	if err != nil {
		return nil, err
	}
	corpAccessToken := corpAccessTokenResp.AccessToken

	opt := &ExternalContactUserInfo{}
	err = e.workWechat.Scan(context.Background(), NewGetExternalContactUserInfo(corpAccessToken, externalUserId, cursor), opt)
	if err != nil {
		return nil, err
	}
	if opt.ErrCode != 0 {
		return nil, errors.New("设置授权配置失败")
	}
	return opt, nil
}

// 发送新客户欢迎语
func NewSendWelcomeMsg(corpAccessToken string, welcomeCode string, text string, attachments []Attachments) Action {
	reqUrl := BaseWeWorkUrl + fmt.Sprintf("/cgi-bin/externalcontact/send_welcome_msg?access_token=%s", corpAccessToken)
	return NewWeWordApi(reqUrl,
		WitchMethod(HttpPost),
		WitchBody(func() (bytes []byte, e error) {
			reqInfo := reqSendWelcomeMsg{
				WelcomeCode: welcomeCode,
				Text: Text{
					Content: text,
				},
				Attachments: attachments,
			}
			jsonInfo, err := json.Marshal(reqInfo)
			if err != nil {
				return nil, err
			}
			return jsonInfo, nil
		}),
	)
}

/**
 * @Description: 发送新客户欢迎语
 * @author:ljj
 * @receiver w
 * @param welcomeCode string 通过添加外部联系人事件推送给企业的发送欢迎语的凭证，有效期为20秒
 * @param text string 消息文本内容,最长为4000字节
 * @param attachments []Attachments 附件，最多可添加9个附件
 * @return *RespSendWelcomeMsg
 * @return error
 */
func (e *externalContact) SendWelcomeMsg(welcomeCode string, text string, attachments []Attachments) (*RespSendWelcomeMsg, error) {
	corpAccessTokenResp, err := e.workWechat.GetCorpAccessToken()
	if err != nil {
		return nil, err
	}
	corpAccessToken := corpAccessTokenResp.AccessToken

	opt := &RespSendWelcomeMsg{}
	err = e.workWechat.Scan(context.Background(), NewSendWelcomeMsg(
		corpAccessToken,
		welcomeCode,
		text,
		attachments,
	), opt)
	if err != nil {
		return nil, err
	}
	if opt.ErrCode != 0 {
		return nil, errors.New("设置授权配置失败")
	}
	return opt, nil
}
