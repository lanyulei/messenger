package common

import (
	"fmt"
	"time"

	"github.com/lanyulei/messenger/config"

	"github.com/lanyulei/messenger/sender/wecom"
	"github.com/lanyulei/messenger/types"

	"github.com/guonaihong/gout"
)

/*
  @Author : lanyulei
  @Desc :
*/

var (
	accessToken map[string]interface{}
)

// GetAccountToken
// @Description: get wecom account token
// @return err
func GetAccountToken() (at string, err error) {
	var (
		result map[string]interface{}
	)

	if accessToken == nil || time.Now().Unix() > accessToken["expires_time"].(int64) {
		if accessToken == nil {
			accessToken = make(map[string]interface{})
		}

		err = gout.GET(wecom.GetAccountTokenURL).
			SetQuery(gout.H{
				"corpid":     config.GetConfig().WeCom.CorpId,
				"corpsecret": config.GetConfig().WeCom.CorpSecret,
			}).
			BindJSON(&result).
			Do()
		if err != nil {
			err = fmt.Errorf("failed to get access token, err:%s", err.Error())
			return
		}

		if errCode, ok := result["errcode"]; ok && int(errCode.(float64)) != 0 {
			err = fmt.Errorf("failed to get wecom access token, err:%s", result["errmsg"])
			return
		}

		accessToken["expires_time"] = time.Now().Unix() + int64(result["expires_in"].(float64))
		accessToken["access_token"] = result["access_token"]
	}

	at = accessToken["access_token"].(string)

	return
}

func GetWeComUserId(mobile string) (result string, err error) {
	var (
		at  string
		req map[string]interface{}
	)

	at, err = GetAccountToken()
	if err != nil {
		err = fmt.Errorf("failed to get wecom account token, err:%s", err.Error())
		return
	}

	err = gout.POST(wecom.GetUserIdURL).
		SetHeader(gout.H{"Content-Type": "application/json"}).
		SetQuery(gout.H{"access_token": at}).
		SetJSON(map[string]interface{}{
			"mobile": mobile,
		}).
		BindJSON(&req).
		Do()
	if err != nil {
		err = fmt.Errorf("failed to get wecom user id, err:%s", err.Error())
		return
	}

	if int(req["errcode"].(float64)) != 0 {
		err = fmt.Errorf("failed to get wecom user id, err:%s", req["errmsg"])
		return
	}

	result = req["userid"].(string)
	return
}

func FormatMarkdown(title string, message *types.Message) (content string) {
	return fmt.Sprintf(`### %s
><font color="comment">标题:</font> %s
><font color="comment">优先级:</font> %s
><font color="comment">申请人:</font> %s
><font color="comment">申请时间:</font> %s
><font color="comment">最近处理时间:</font> %s`,
		title,
		message.Title,
		message.Priority,
		message.Creator,
		message.CreatedAt,
		message.UpdatedAt,
	)
}
