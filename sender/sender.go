package sender

import (
	dingtalkNotification "github.com/lanyulei/messenger/sender/dingtalk/notification"
	"github.com/lanyulei/messenger/sender/email"
	larkNotification "github.com/lanyulei/messenger/sender/lark/notification"
	"github.com/lanyulei/messenger/sender/webhook"
	wecomNotification "github.com/lanyulei/messenger/sender/wecom/notification"
	"github.com/lanyulei/messenger/types"
)

/*
  @Author : lanyulei
  @Desc :
*/

type Interface interface {
	Email(to, cc []string, title string, content *types.Message) error
	DingTalkNotify(content map[string]interface{}, userIdList, deptIdList string, toAllUser bool) (result string, err error)
	WeComNotify(content map[string]interface{}) (result map[string]interface{}, err error)
	LarkNotify(mobiles []string, content map[string]interface{}) (result map[string]interface{}, err error)
	Webhook(content map[string]interface{}, webhook []string) error
}

type sender struct{}

func New() Interface {
	return &sender{}
}

// Email
// @Description: send email notifications
func (n *sender) Email(to, cc []string, title string, content *types.Message) (err error) {
	return email.Send(to, cc, title, content)
}

// DingTalkNotify
// @Description: send DingTalk work notification message
func (n *sender) DingTalkNotify(content map[string]interface{}, userIdList, deptIdList string, toAllUser bool) (result string, err error) {
	return dingtalkNotification.Send(userIdList, deptIdList, toAllUser, content)
}

// LarkNotify
// @Description: send lark notification
func (n *sender) LarkNotify(mobiles []string, content map[string]interface{}) (result map[string]interface{}, err error) {
	return larkNotification.Send(mobiles, content)
}

// WeComNotify
// @Description: send enterprise WeChat messages
func (n *sender) WeComNotify(content map[string]interface{}) (result map[string]interface{}, err error) {
	return wecomNotification.Send(content)
}

// Webhook
// @Description: send webhook notice
func (n *sender) Webhook(content map[string]interface{}, webhookList []string) (err error) {
	return webhook.Send(webhookList, content)
}
