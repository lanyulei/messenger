package dingtalk

/*
  @Author : lanyulei
  @Desc :
*/

const (
	NotifyBaseURL     = "https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2"
	NotifyResultURL   = "https://oapi.dingtalk.com/topapi/message/corpconversation/getsendresult"
	GetAccessTokenURL = "https://oapi.dingtalk.com/gettoken"
	GetUserIdURL      = "https://oapi.dingtalk.com/topapi/v2/user/getbymobile"

	MsgTextType     = "text"
	MsgImageType    = "image"
	MsgVoiceType    = "voice"
	MsgFileType     = "file"
	MsgLinkType     = "link"
	MsgMarkdownType = "markdown"
)

type UserIdResponse struct {
	RequestID string `json:"request_id"`
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	Result    struct {
		UserID                     string   `json:"userid"`
		ExclusiveAccountUserIDList []string `json:"exclusive_account_userid_list"`
	} `json:"result"`
}
