package config

/*
  @Author : lanyulei
  @Desc :
*/

var configuration Config

type Email struct {
	Alias    string `json:"alias"`    // sender alias
	Host     string `json:"host"`     // smtp host
	Port     int    `json:"port"`     // smtp port
	User     string `json:"user"`     // sender email
	Password string `json:"password"` // sender password
}

type DingTalk struct {
	AgentID   string `json:"agent_id"`   // agent id
	AppKey    string `json:"app_key"`    // app key
	AppSecret string `json:"app_secret"` // app secret
}

type WeCom struct {
	AgentId    string `json:"agent_id"`    // agent id
	CorpId     string `json:"corp_id"`     // corp id
	CorpSecret string `json:"corp_secret"` // corp secret
}

type Lark struct {
	AppId     string `json:"app_id"`     // app id
	AppSecret string `json:"app_secret"` // app secret
}

// Config
// @Description: configure the required format
type Config struct {
	Email    Email    `json:"email"`
	DingTalk DingTalk `json:"dingtalk"`
	WeCom    WeCom    `json:"wecom"`
	Lark     Lark     `json:"lark"`
}

const (
	MessengerEmailAlias    = "MESSENGER_EMAIL_ALIAS"
	MessengerEmailHost     = "MESSENGER_EMAIL_HOST"
	MessengerEmailPort     = "MESSENGER_EMAIL_PORT"
	MessengerEmailUser     = "MESSENGER_EMAIL_USER"
	MessengerEmailPassword = "MESSENGER_EMAIL_PASSWORD"

	MessengerDingTalkAgentID   = "MESSENGER_DINGTALK_AGENT_ID"
	MessengerDingTalkAppKey    = "MESSENGER_DINGTALK_APP_KEY"
	MessengerDingTalkAppSecret = "MESSENGER_DINGTALK_APP_SECRET"

	MessengerWeComAgentId    = "MESSENGER_WECOM_AGENT_ID"
	MessengerWeComCorpId     = "MESSENGER_WECOM_CORP_ID"
	MessengerWeComCorpSecret = "MESSENGER_WECOM_CORP_SECRET"

	MessengerLarkAppId   = "MESSENGER_LARK_APP_ID"
	MessageLarkAppSecret = "MESSENGER_LARK_APP_SECRET"
)
