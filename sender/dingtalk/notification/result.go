package notification

import (
	"encoding/json"
	"fmt"
	"github.com/lanyulei/messenger/config"
	"github.com/lanyulei/messenger/sender/dingtalk"
	"github.com/lanyulei/messenger/sender/dingtalk/common"

	"github.com/guonaihong/gout"
)

/*
  @Author : lanyulei
  @Desc :
*/

func GetResult(taskId int) (res []byte, err error) {
	var (
		data        []byte
		accessToken string
	)

	accessToken, err = common.GetAccountToken()
	if err != nil {
		return
	}

	params := map[string]interface{}{
		"agent_id": config.GetConfig().DingTalk.AgentID,
		"task_id":  taskId,
	}

	data, err = json.Marshal(params)
	if err != nil {
		err = fmt.Errorf("json serialization failed, %s", err.Error())
		return
	}

	err = gout.POST(dingtalk.NotifyResultURL).
		SetHeader(gout.H{"Content-Type": "application/json"}).
		SetQuery(gout.H{"access_token": accessToken}).
		SetBody(data).
		BindBody(&res).
		Do()
	if err != nil {
		err = fmt.Errorf("failed to get notification details by task id, %s", err.Error())
		return
	}

	return
}
