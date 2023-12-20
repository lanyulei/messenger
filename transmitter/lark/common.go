package lark

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

// GetLarkAccountToken
// @Description: get lark tenant account token
// @return err
func GetLarkAccountToken() (at string, err error) {
	var (
		result map[string]interface{}
	)

	if accessToken == nil || time.Now().Unix() > accessToken["expires_time"].(int64) {
		if accessToken == nil {
			accessToken = make(map[string]interface{})
		}
		err = gout.POST(GetTenantAccountTokenURL).SetJSON(gout.H{
			"app_id":     config.GetConfig().Lark.AppId,
			"app_secret": config.GetConfig().Lark.AppSecret,
		}).BindJSON(&result).Do()
		if err != nil {
			err = fmt.Errorf("failed to get access token, err:%s", err.Error())
			return
		}

		if code, ok := result["code"]; !ok || int(code.(float64)) != 0 {
			err = fmt.Errorf("failed to get lark access token, err:%s", result["msg"])
			return
		}

		accessToken["expires_time"] = time.Now().Unix() + int64(result["expire"].(float64))
		accessToken["access_token"] = result["tenant_access_token"]
	}

	at = accessToken["access_token"].(string)

	return
}

// GetLarkUserIDByMobiles
// @Description: get userId of lark user by cell phone number
// @param mobiles
// @return larkUserResponse
// @return err
func GetLarkUserIDByMobiles(mobiles []string) (larkUserResponse map[string]interface{}, err error) {
	var (
		at string
	)

	at, err = GetLarkAccountToken()
	if err != nil {
		err = fmt.Errorf("failed to get lark account token, err:%s", err.Error())
		return
	}

	err = gout.POST(GetLarkUserIDByMobilesURL).
		SetQuery(gout.H{"user_id_type": "user_id"}).
		SetHeader(gout.H{"Content-Type": "application/json", "Authorization": "Bearer " + at}).
		SetJSON(gout.H{"mobiles": mobiles}).
		BindJSON(&larkUserResponse).
		Do()
	if err != nil {
		err = fmt.Errorf("failed to get lark user id by mobiles, err:%s", err.Error())
		return
	}

	if int(larkUserResponse["code"].(float64)) != 0 {
		err = fmt.Errorf("failed to get lark user id by mobiles, err:%s", larkUserResponse["msg"].(string))
		return
	}

	return
}
