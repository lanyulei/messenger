package common

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/guonaihong/gout"
	"github.com/lanyulei/messenger/config"
	"github.com/lanyulei/messenger/sender/dingtalk"
	"github.com/lanyulei/messenger/types"
)

/*
  @Author : lanyulei
  @Desc :
*/

var (
	accessToken map[string]interface{}
)

// GetAccountToken
// @Description: get dingtalk account token
// @return err
func GetAccountToken() (at string, err error) {
	var (
		result []byte
	)

	if accessToken == nil || time.Now().Unix() > accessToken["expires_time"].(int64) {
		err = gout.GET(dingtalk.GetAccessTokenURL).SetQuery(gout.H{
			"appkey":    config.GetConfig().DingTalk.AppKey,
			"appsecret": config.GetConfig().DingTalk.AppSecret,
		}).BindBody(&result).Do()
		if err != nil {
			err = fmt.Errorf("failed to get access token, err:%s", err.Error())
			return
		}

		err = json.Unmarshal(result, &accessToken)
		if err != nil {
			err = fmt.Errorf("failed to get access token, err:%s", err.Error())
			return
		}

		if errCode, ok := accessToken["errcode"]; ok && int(errCode.(float64)) != 0 {
			err = fmt.Errorf("failed to get dingtalk access token, err:%s", accessToken["errmsg"].(string))
			return
		}

		accessToken["expires_time"] = time.Now().Unix() + int64(accessToken["expires_in"].(float64))
	}

	at = accessToken["access_token"].(string)

	return
}

// GetDingtalkUserId
// @Description: obtain the user ID of DingTalk users through their phone number, and enterprise applications need to have this interface permission, https://open.dingtalk.com/document/orgapp/obtain-the-userid-of-your-mobile-phone-number
// @param mobile phone number
// @return res return results
// @return err error message
func GetDingtalkUserId(mobile string) (res dingtalk.UserIdResponse, err error) {
	var (
		at string
	)

	at, err = GetAccountToken()
	if err != nil {
		err = fmt.Errorf("failed to get account token, err:%s", err.Error())
		return
	}

	err = gout.POST(dingtalk.GetUserIdURL).
		SetQuery(gout.H{
			"access_token": at,
		}).
		SetHeader(gout.H{
			"Accept": "application/json",
		}).
		SetJSON(gout.H{
			"mobile": mobile,
		}).
		BindJSON(&res).Do()
	if err != nil {
		err = fmt.Errorf("failed to get user id, err:%s", err.Error())
		return
	}

	return
}

func FormatMarkdown(title string, message *types.Message) (res string) {
	return fmt.Sprintf("### %s  \n  > 标题：%s  \n  > 优先级：%s  \n  > 申请人：%s  \n  > 申请时间：%s  \n  > 最近处理时间：%s",
		title,
		message.Title,
		message.Priority,
		message.Creator,
		message.CreatedAt,
		message.UpdatedAt,
	)
}
