package notification

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/guonaihong/gout"
	"github.com/lanyulei/messenger/config"
	"github.com/lanyulei/messenger/sender/dingtalk"
	"github.com/lanyulei/messenger/sender/dingtalk/common"
)

/*
  @Author : lanyulei
  @Desc :
*/

// Send
// @Description: send dingtalk job notification
// @param userIdList comma-separated list of user ids, max list length: 20
// @param deptIdList comma-separated list of department ids, maximum list length: 20
// @param toAllUser whether to send to all users of the enterprise
// @param msg message content
// @return result return result
// @return err
func Send(userIdList, deptIdList string, toAllUser bool, msg map[string]interface{}) (result string, err error) {
	var (
		data, res   []byte
		respMap     map[string]interface{}
		accessToken string
		taskId      int
	)

	accessToken, err = common.GetAccountToken()
	if err != nil {
		return
	}

	body := map[string]interface{}{
		"agent_id":    config.GetConfig().DingTalk.AgentID,
		"to_all_user": toAllUser,
		"msg":         msg,
	}

	if !toAllUser {
		if userIdList == "" && deptIdList == "" {
			err = errors.New("if to_all_user is false, there must be a valid value in userid_list or dept_id_list")
			return
		}
	}

	if userIdList != "" {
		body["userid_list"] = userIdList
	}

	if deptIdList != "" {
		body["dept_id_list"] = deptIdList
	}

	data, err = json.Marshal(body)
	if err != nil {
		err = fmt.Errorf("json serialization failed, %s", err.Error())
		return
	}

	err = gout.POST(dingtalk.NotifyBaseURL).
		SetHeader(gout.H{"Content-Type": "application/json"}).
		SetQuery(gout.H{"access_token": accessToken}).
		SetBody(data).
		BindBody(&res).
		Do()
	if err != nil {
		err = fmt.Errorf("send dingtalk work notification, %s", err.Error())
		return
	}

	err = json.Unmarshal(res, &respMap)
	if err != nil {
		err = fmt.Errorf("json deserialization failed, %s", err.Error())
		return
	}

	if t, ok := respMap["task_id"]; ok {
		taskId = int(t.(float64))
		res, err = GetResult(taskId)
		if err != nil {
			err = fmt.Errorf("get dingtalk work notification result failed, %s, task id: %d", err.Error(), taskId)
			return
		}
		result = fmt.Sprintf("job notification return results, task id %d, %s", taskId, string(res))
	} else {
		err = fmt.Errorf("send dingtalk work notification failed, %s", string(res))
	}

	return
}
