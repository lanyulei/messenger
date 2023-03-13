package notification

import (
    "encoding/json"
    "fmt"
    "github.com/guonaihong/gout"
    "messenger/config"
    "messenger/notify/dingtalk"
)

/*
  @Author : lanyulei
  @Desc :
*/

func GetResult(taskId int) (res []byte, err error) {
    var (
        data         []byte
        accountToken string
    )

    accountToken, err = dingtalk.GetAccountToken()
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
        SetQuery(gout.H{"access_token": accountToken}).
        SetBody(data).
        BindBody(&res).
        Do()
    if err != nil {
        err = fmt.Errorf("failed to get notification details by task id, %s", err.Error())
        return
    }

    return
}
