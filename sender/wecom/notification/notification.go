package notification

import (
	"fmt"

	"github.com/guonaihong/gout"
	"github.com/lanyulei/messenger/sender/wecom"
	"github.com/lanyulei/messenger/sender/wecom/common"
)

/*
  @Author : lanyulei
  @Desc :
*/

func Send(content map[string]interface{}) (result map[string]interface{}, err error) {
	var (
		at string
	)

	at, err = common.GetAccountToken()
	if err != nil {
		err = fmt.Errorf("failed to get access token, err:%s", err.Error())
		return
	}

	err = gout.POST(wecom.SendMessageURL).
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
