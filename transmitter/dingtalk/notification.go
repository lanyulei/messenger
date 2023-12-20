package dingtalk

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/guonaihong/gout"
	"github.com/lanyulei/messenger/config"
)

/*
  @Author : lanyulei
  @Desc :
*/

type Interface interface {
	Send(msg map[string]interface{}) (result string, err error)
	SendText(content string) (result string, err error)
	SendImage(mediaId string) (result string, err error)
	SendVoice(mediaId, duration string) (result string, err error)
	SendFile(mediaId string) (result string, err error)
	SendLink(title, text, messageUrl, picUrl string) (result string, err error)
	SendMarkdown(title, text string) (result string, err error)
}

type handler struct {
	UserIdList string
	DeptIdList string
	ToAllUser  bool
}

func New(userIdList, deptIdList string, toAllUser bool) Interface {
	return &handler{
		UserIdList: userIdList,
		DeptIdList: deptIdList,
		ToAllUser:  toAllUser,
	}
}

// Send OA messages and card messages are sent using this method
// The message structure is configured according to this document, https://open.dingtalk.com/document/orgapp/message-types-and-data-format#title-8w7-rsl-obh
func (h *handler) Send(msg map[string]interface{}) (result string, err error) {
	var (
		data, resValue []byte
		respMap, res   map[string]interface{}
		at             string
		taskId         int
	)

	at, err = GetAccountToken()
	if err != nil {
		return
	}

	body := map[string]interface{}{
		"agent_id":    config.GetConfig().DingTalk.AgentID,
		"to_all_user": h.ToAllUser,
		"msg":         msg,
	}

	if !h.ToAllUser {
		if h.UserIdList == "" && h.DeptIdList == "" {
			err = errors.New("if to_all_user is false, there must be a valid value in userid_list or dept_id_list")
			return
		}
	}

	if h.UserIdList != "" {
		body["userid_list"] = h.UserIdList
	}

	if h.DeptIdList != "" {
		body["dept_id_list"] = h.DeptIdList
	}

	data, err = json.Marshal(body)
	if err != nil {
		err = fmt.Errorf("json serialization failed, %s", err.Error())
		return
	}

	err = gout.POST(NotifyBaseURL).
		SetHeader(gout.H{"Content-Type": "application/json"}).
		SetQuery(gout.H{"access_token": at}).
		SetBody(data).
		BindJSON(&respMap).
		Do()
	if err != nil {
		err = fmt.Errorf("send dingtalk work notification, %s", err.Error())
		return
	}

	if int(respMap["errcode"].(float64)) != 0 {
		err = fmt.Errorf("send dingtalk work notification failed, %s", respMap["errmsg"].(string))
		return
	}

	if t, ok := respMap["task_id"]; ok {
		taskId = int(t.(float64))
		res, err = GetResult(taskId)
		if err != nil {
			err = fmt.Errorf("get dingtalk work notification result failed, %s, task id: %d", err.Error(), taskId)
			return
		}
		resValue, err = json.Marshal(res)
		if err != nil {
			err = fmt.Errorf("json serialization failed, %s", err.Error())
			return
		}
		result = fmt.Sprintf("job notification return results, task id %d, %s", taskId, string(resValue))
	} else {
		err = fmt.Errorf("send dingtalk work notification failed, %s", string(resValue))
	}

	return
}

// SendText send text message
func (h *handler) SendText(content string) (result string, err error) {
	msg := map[string]interface{}{
		"msgtype": MsgTextType,
		"text": map[string]interface{}{
			"content": content,
		},
	}
	return h.Send(msg)
}

// SendImage send image message
func (h *handler) SendImage(mediaId string) (result string, err error) {
	msg := map[string]interface{}{
		"msgtype": MsgImageType,
		"image": map[string]interface{}{
			"media_id": mediaId,
		},
	}
	return h.Send(msg)
}

// SendVoice send a voice message
func (h *handler) SendVoice(mediaId, duration string) (result string, err error) {
	msg := map[string]interface{}{
		"msgtype": MsgVoiceType,
		"voice": map[string]interface{}{
			"media_id": mediaId,
			"duration": duration,
		},
	}
	return h.Send(msg)
}

// SendFile send a file message
func (h *handler) SendFile(mediaId string) (result string, err error) {
	msg := map[string]interface{}{
		"msgtype": MsgFileType,
		"file": map[string]interface{}{
			"media_id": mediaId,
		},
	}
	return h.Send(msg)
}

// SendLink send a link message
func (h *handler) SendLink(title, text, messageUrl, picUrl string) (result string, err error) {
	msg := map[string]interface{}{
		"msgtype": MsgLinkType,
		"link": map[string]interface{}{
			"title":      title,
			"text":       text,
			"messageUrl": messageUrl,
			"picUrl":     picUrl,
		},
	}
	return h.Send(msg)
}

// SendMarkdown send a markdown message
func (h *handler) SendMarkdown(title, text string) (result string, err error) {
	msg := map[string]interface{}{
		"msgtype": MsgMarkdownType,
		"markdown": map[string]interface{}{
			"title": title,
			"text":  text,
		},
	}
	return h.Send(msg)
}
