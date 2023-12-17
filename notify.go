package notify

import (
	"errors"
	"strings"

	"github.com/lanyulei/messenger/sender"
	"github.com/lanyulei/messenger/types"
)

/*
  @Author : lanyulei
  @Desc :
*/

func Notify(transmitter string, to []string, content interface{}, option map[string]interface{}) (err error) {
	s := sender.New()

	switch transmitter {
	case types.Email:
		var (
			title string
			cc    []string
		)
		if _, ok := option["cc"]; ok {
			cc = option["cc"].([]string)
		}
		if _, ok := option["title"]; ok {
			title = option["title"].(string)
		}
		err = s.Email(to, cc, title, content.(string)) // content is text/html
		if err != nil {
			err = errors.New("failed to send email")
		}
	case types.DingTalk:
		var (
			toAllUser  bool
			deptIdList []string
		)
		if _, ok := option["toAllUser"]; ok {
			toAllUser = option["toAllUser"].(bool)
		}
		if _, ok := option["deptIdList"]; ok {
			deptIdList = option["deptIdList"].([]string)
		}
		_, err = s.DingTalkNotify(content.(map[string]interface{}), strings.Join(to, "|"), strings.Join(deptIdList, "|"), toAllUser)
		if err != nil {
			err = errors.New("failed to send dingtalk")
		}
	case types.Lark:
		_, err = s.LarkNotify(to, content.(map[string]interface{}))
		if err != nil {
			err = errors.New("failed to send lark")
		}
	case types.WeCom:
		_, err = s.WeComNotify(to, content.(map[string]interface{}))
		if err != nil {
			err = errors.New("failed to send wecom")
		}
	case types.Webhook:
		err = s.Webhook(to, content.(map[string]interface{}))
		if err != nil {
			err = errors.New("failed to send webhook")
		}
	default:
		err = errors.New("unsupported sending type, please confirm")
	}

	return
}
