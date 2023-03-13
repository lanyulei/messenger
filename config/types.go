package config

/*
  @Author : lanyulei
  @Desc :
*/

var configuration Config

// Config
// @Description: configure the required format
type Config struct {
    Email struct {
        Alias    string `json:"alias"`    // sender alias
        Host     string `json:"host"`     // smtp host
        Port     int    `json:"port"`     // smtp port
        User     string `json:"user"`     // sender email
        Password string `json:"password"` // sender password
    } `json:"email"`
    DingTalk struct {
        AgentID   string `json:"agent_id"`   // agent id
        AppKey    string `json:"app_key"`    // app key
        AppSecret string `json:"app_secret"` // app secret
    } `json:"dingtalk"`
    WeCom struct {
        AgentID   string `json:"agent_id"`   // agent id
        AppKey    string `json:"app_key"`    // app key
        AppSecret string `json:"app_secret"` // app secret
    } `json:"wecom"`
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

    MessengerWeComAgentID   = "MESSENGER_WECOM_AGENT_ID"
    MessengerWeComAppKey    = "MESSENGER_WECOM_APP_KEY"
    MessengerWeComAppSecret = "MESSENGER_WECOM_APP_SECRET"
)
