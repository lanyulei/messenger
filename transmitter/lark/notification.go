package lark

import (
	"fmt"

	"github.com/guonaihong/gout"
)

/*
  @Author : lanyulei
  @Desc :
*/

type Interface interface {
	Send(msg map[string]interface{}) (result map[string]interface{}, err error)
	SendText(content string) (result map[string]interface{}, err error)
	SendImage(mediaId string) (result map[string]interface{}, err error)
	SendShareChat(shareChatId string) (result map[string]interface{}, err error)
}

type handler struct {
	Mobiles []string
}

func New(mobiles []string) Interface {
	return &handler{
		Mobiles: mobiles,
	}
}

// Send notification
// more ways to send, https://open.feishu.cn/document/server-docs/im-v1/batch_message/send-messages-in-batches
func (h *handler) Send(msg map[string]interface{}) (result map[string]interface{}, err error) {
	var (
		at                  string
		receiveIds          []string
		larkUserIdsResponse map[string]interface{}
	)

	// get account token
	at, err = GetLarkAccountToken()
	if err != nil {
		err = fmt.Errorf("get account token failed, %s", err.Error())
		return
	}

	// get user id
	larkUserIdsResponse, err = GetLarkUserIDByMobiles(h.Mobiles)
	if err != nil {
		err = fmt.Errorf("get user id failed, %s", err.Error())
		return
	}

	if d, ok := larkUserIdsResponse["data"]; ok {
		if u, ok := d.(map[string]interface{})["user_list"]; ok {
			for _, v := range u.([]interface{}) {
				if userId, ok := v.(map[string]interface{})["user_id"]; ok {
					receiveIds = append(receiveIds, userId.(string))
				}
			}
		}
	}

	if _, ok := msg["user_ids"]; ok && msg["user_ids"] != nil && len(msg["user_ids"].([]string)) > 0 {
		receiveIds = append(receiveIds, msg["user_ids"].([]string)...)
	}

	msg["user_ids"] = receiveIds

	err = gout.POST(NotifyBaseURL).
		SetHeader(gout.H{"Content-Type": "application/json", "Authorization": "Bearer " + at}).
		SetJSON(msg).
		BindJSON(&result).
		Do()
	if err != nil {
		err = fmt.Errorf("send notification failed, %s", err.Error())
		return
	}
	if code, ok := result["code"]; ok && int(code.(float64)) == 0 {
		//logger.Infof("send notification success, result: %s", result["msg"])
	} else {
		err = fmt.Errorf("send notification failed, result: %s", result["msg"])
	}

	return
}

func (h *handler) SendText(content string) (result map[string]interface{}, err error) {
	msg := map[string]interface{}{
		"msg_type": MessageTextType,
		"content": map[string]interface{}{
			"text": content,
		},
	}
	return h.Send(msg)
}

func (h *handler) SendImage(mediaId string) (result map[string]interface{}, err error) {
	msg := map[string]interface{}{
		"msg_type": MessageImageType,
		"content": map[string]interface{}{
			"image_key": mediaId,
		},
	}
	return h.Send(msg)
}

func (h *handler) SendShareChat(shareChatId string) (result map[string]interface{}, err error) {
	msg := map[string]interface{}{
		"msg_type": MessageShareType,
		"content": map[string]interface{}{
			"share_chat_id": shareChatId,
		},
	}
	return h.Send(msg)
}
