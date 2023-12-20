package dingtalk

import (
	"fmt"
	"time"

	"github.com/guonaihong/gout"
	"github.com/lanyulei/messenger/config"
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
		result map[string]interface{}
	)

	if accessToken == nil || time.Now().Unix() > accessToken["expires_time"].(int64) {
		if accessToken == nil {
			accessToken = make(map[string]interface{})
		}

		err = gout.GET(GetAccessTokenURL).SetQuery(gout.H{
			"appkey":    config.GetConfig().DingTalk.AppKey,
			"appsecret": config.GetConfig().DingTalk.AppSecret,
		}).BindJSON(&result).Do()
		if err != nil {
			err = fmt.Errorf("failed to get access token, err:%s", err.Error())
			return
		}

		if errCode, ok := result["errcode"]; ok && int(errCode.(float64)) != 0 {
			err = fmt.Errorf("failed to get dingtalk access token, err:%s", result["errmsg"].(string))
			return
		}

		accessToken["expires_time"] = time.Now().Unix() + int64(result["expires_in"].(float64))
		accessToken["access_token"] = result["access_token"]
	}

	at = accessToken["access_token"].(string)

	return
}

// GetDingtalkUserId
// @Description: obtain the user ID of DingTalk users through their phone number, and enterprise applications need to have this interface permission, https://open.dingtalk.com/document/orgapp/obtain-the-userid-of-your-mobile-phone-number
// @param mobile phone number
// @return res return results
// @return err error message
func GetDingtalkUserId(mobile string) (res UserIdResponse, err error) {
	var (
		at string
	)

	at, err = GetAccountToken()
	if err != nil {
		err = fmt.Errorf("failed to get account token, err:%s", err.Error())
		return
	}

	// requires qyapi_get_member_by_mobile permission
	err = gout.POST(GetUserIdURL).
		SetQuery(gout.H{"access_token": at}).
		SetHeader(gout.H{"Accept": "application/json"}).
		SetJSON(gout.H{"mobile": mobile}).
		BindJSON(&res).Do()
	if err != nil {
		err = fmt.Errorf("failed to get user id, err:%s", err.Error())
		return
	}

	if res.ErrCode != 0 {
		err = fmt.Errorf("failed to get user id, err: %s", res.ErrMsg)
		return
	}

	return
}
