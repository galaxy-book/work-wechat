package work

type SendMsgResp struct {
	respCommon
	InvalidUser  string `json:"invaliduser"`
	InvalidParty string `json:"invalidparty"`
	InvalidTag   string `json:"invalidtag"`
}

type SendMsgReq struct {
	ToUser                 string                  `json:"touser"`
	ToParty                string                  `json:"toparty"`
	ToTag                  string                  `json:"totag"`
	MsgType                string                  `json:"msgtype"`
	AgentId                int64                   `json:"agentid"`
	Safe                   int                     `json:"safe"`
	Text                   *MsgText                `json:"text"`
	Image                  *MsgImage               `json:"image"`
	Voice                  *MsgVoice               `json:"voice"`
	Video                  *MsgVideo               `json:"video"`
	File                   *MsgFile                `json:"file"`
	TextCard               *MsgTextCard            `json:"textcard"`
	News                   *MsgNews                `json:"news"`
	MpNews                 *MsgMpNews              `json:"mpnews"`
	Markdown               *MsgMarkdown            `json:"markdown"`
	InteractvieTaskCard    *MsgInteractiveTaskCard `json:"interactive_taskcard"`
	EnableIdTrans          int                     `json:"enable_id_trans"`
	EnableDuplicateCheck   int                     `json:"enable_duplicate_check"`
	DuplicateCheckInterval int                     `json:"duplicate_check_interval"`
}

type MsgText struct {
	Content string `json:"content"`
}

type MsgImage struct {
	MediaId string `json:"mediaId"`
}

type MsgVoice struct {
	MediaId string `json:"mediaId"`
}

type MsgVideo struct {
	MediaId     string `json:"mediaId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type MsgFile struct {
	MediaId string `json:"media_id"`
}

type MsgTextCard struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

type MsgNews struct {
	Articles MsgNewsArticles `json:"articles"`
}

type MsgNewsArticles struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Picurl      string `json:"picurl"`
}

type MsgMpNews struct {
	Articles MsgMpNewsArticles `json:"articles"`
}

type MsgMpNewsArticles struct {
	Title            string `json:"title"`
	ThumbMediaId     string `json:"thumb_media_id"`
	Author           string `json:"author"`
	ContentSourceUrl string `json:"content_source_url"`
	Content          string `json:"content"`
	Digest           string `json:"digest"`
}

type MsgMarkdown struct {
	Content string `json:"content"`
}

type MsgInteractiveTaskCard struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Url         string   `json:"url"`
	TaskId      string   `json:"task_id"`
	Btn         []MsgBtn `json:"btn"`
}

type MsgBtn struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	Color  string `json:"color"`
	IsBold bool   `json:"is_bold"`
}
