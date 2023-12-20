package dingtalk

import (
	"encoding/json"
	"fmt"

	"github.com/lanyulei/messenger/config"

	"github.com/guonaihong/gout"
)

/*
  @Author : lanyulei
  @Desc :
*/

func GetResult(taskId int) (resMap map[string]interface{}, err error) {
	var (
		data []byte
		at   string
	)

	at, err = GetAccountToken()
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

	err = gout.POST(NotifyResultURL).
		SetHeader(gout.H{"Content-Type": "application/json"}).
		SetQuery(gout.H{"access_token": at}).
		SetBody(data).
		BindJSON(&resMap).
		Do()
	if err != nil {
		err = fmt.Errorf("failed to get notification details by task id, %s", err.Error())
		return
	}

	if int(resMap["errcode"].(float64)) != 0 {
		err = fmt.Errorf("failed to get notification details by task id, %s", resMap["errmsg"].(string))
		return
	}

	return
}
