package wecom

import (
	"fmt"

	"github.com/guonaihong/gout"
)

/*
  @Author : lanyulei
  @Desc :
*/

type Interface interface {
	Send(content map[string]interface{}) (result map[string]interface{}, err error)
}

type handler struct {
	ToUser                 string // Specify the member to receive the message, a list of member IDs (multiple recipients are separated by ‘|’, up to 1000 are supported). Special case: If specified as "@all", it will be sent to all members of the enterprise application.
	ToParty                string // Specify the department to receive the message, department ID list, multiple recipients separated by ‘|’, up to 100 are supported. This parameter is ignored when touser is "@all"
	ToTag                  string // Specify the tag to receive the message, tag ID list, multiple recipients separated by ‘|’, up to 100 are supported. This parameter is ignored when touser is "@all"
	AgentId                int    // The id of the enterprise application, integer type. Can be viewed on the app’s settings page
	Safe                   int    // Indicates whether the message is confidential, 0 means no, 1 means yes, the default is 0
	EnableIdTrans          int    // Indicates whether to enable id translation; 0 means no, 1 means yes, the default is 0
	EnableDuplicateCheck   int    // Indicates whether to enable duplicate message checking, 0 means no, 1 means yes, the default is 0
	DuplicateCheckInterval int    // Indicates the time interval for repeated message checking. The default is 1800 s and the maximum is no more than 4 hours.
}

func New(agentId int, options map[string]interface{}) Interface {
	h := &handler{
		AgentId: agentId,
	}

	if v, ok := options["touser"]; ok {
		h.ToUser = v.(string)
	}

	if v, ok := options["toparty"]; ok {
		h.ToParty = v.(string)
	}

	if v, ok := options["totag"]; ok {
		h.ToTag = v.(string)
	}

	if v, ok := options["safe"]; ok {
		h.Safe = v.(int)
	}

	if v, ok := options["enable_id_trans"]; ok {
		h.EnableIdTrans = v.(int)
	}

	if v, ok := options["enable_duplicate_check"]; ok {
		h.EnableDuplicateCheck = v.(int)
	}

	if v, ok := options["duplicate_check_interval"]; ok {
		h.DuplicateCheckInterval = v.(int)
	} else {
		h.DuplicateCheckInterval = 1800
	}

	return h
}

func (h *handler) Send(content map[string]interface{}) (result map[string]interface{}, err error) {
	var (
		at string
	)

	at, err = GetAccountToken()
	if err != nil {
		err = fmt.Errorf("failed to get access token, err:%s", err.Error())
		return
	}

	if h.ToUser != "" {
		if _, ok := content["touser"]; !ok {
			content["touser"] = h.ToUser
		} else {
			content["touser"] = content["touser"].(string) + "|" + h.ToUser
		}
	}

	if h.ToParty != "" {
		if _, ok := content["toparty"]; !ok {
			content["toparty"] = h.ToParty
		} else {
			content["toparty"] = content["toparty"].(string) + "|" + h.ToParty
		}
	}

	if h.ToTag != "" {
		if _, ok := content["totag"]; !ok {
			content["totag"] = h.ToTag
		} else {
			content["totag"] = content["totag"].(string) + "|" + h.ToTag
		}
	}

	if _, ok := content["agentid"]; !ok {
		content["agentid"] = h.AgentId
	}

	if _, ok := content["safe"]; !ok {
		content["safe"] = h.Safe
	}

	if _, ok := content["enable_id_trans"]; !ok {
		content["enable_id_trans"] = h.EnableIdTrans
	}

	if _, ok := content["enable_duplicate_check"]; !ok {
		content["enable_duplicate_check"] = h.EnableDuplicateCheck
	}

	if _, ok := content["duplicate_check_interval"]; !ok {
		content["duplicate_check_interval"] = h.DuplicateCheckInterval
	}

	err = gout.POST(SendMessageURL).
		SetHeader(gout.H{"Content-Type": "application/json"}).
		SetQuery(gout.H{"access_token": at}).
		SetJSON(content).
		BindJSON(&result).
		Do()
	if err != nil {
		err = fmt.Errorf("failed to send message, err:%s", err.Error())
		return
	}

	if int(result["errcode"].(float64)) != 0 {
		err = fmt.Errorf("failed to send wecom message, errcode:%d, errmsg:%s", int(result["errcode"].(float64)), result["errmsg"].(string))
		return
	}

	return
}

func (h *handler) SendText(content string) (result map[string]interface{}, err error) {
	msg := map[string]interface{}{
		"msgtype": MessageTextType,
		"text": map[string]interface{}{
			"content": content,
		},
	}
	return h.Send(msg)
}
