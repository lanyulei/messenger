package lark

/*
  @Author : lanyulei
  @Desc :
*/

const (
	NotifyBaseURL             = "https://open.feishu.cn/open-apis/message/v4/batch_send"
	GetTenantAccountTokenURL  = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"
	GetLarkUserIDByMobilesURL = "https://open.feishu.cn/open-apis/contact/v3/users/batch_get_id"

	MessageCardType  = "interactive"
	MessageTextType  = "text"
	MessageImageType = "image"
	MessagePostType  = "post"
	MessageShareType = "share_chat"
)
