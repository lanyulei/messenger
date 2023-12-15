package common

import (
	"fmt"
	"time"

	"github.com/lanyulei/messenger/config"
	"github.com/lanyulei/messenger/sender/lark"
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

// GetTenantAccountToken
// @Description: get lark tenant account token
// @return err
func GetTenantAccountToken() (at string, err error) {
	var (
		result map[string]interface{}
	)

	if accessToken == nil || time.Now().Unix() > accessToken["expires_time"].(int64) {
		if accessToken == nil {
			accessToken = make(map[string]interface{})
		}
		err = gout.POST(lark.GetTenantAccountTokenURL).SetJSON(gout.H{
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

	at, err = GetTenantAccountToken()
	if err != nil {
		err = fmt.Errorf("failed to get lark account token, err:%s", err.Error())
		return
	}

	err = gout.POST(lark.GetLarkUserIDByMobilesURL).
		SetQuery(gout.H{"user_id_type": "user_id"}).
		SetHeader(gout.H{"Content-Type": "application/json", "Authorization": "Bearer " + at}).
		SetJSON(gout.H{"mobiles": mobiles}).
		BindJSON(&larkUserResponse).
		Do()
	if err != nil {
		err = fmt.Errorf("failed to get lark user id by mobiles, err:%s", err.Error())
		return
	}

	return
}

// CardMessageFormat
// @Description: format card message
// @param title
// @param message
// @return map[string]interface{}
func CardMessageFormat(title string, message *types.Message) map[string]interface{} {
	return map[string]interface{}{
		"header": map[string]interface{}{
			"title": map[string]interface{}{"content": title, "tag": "plain_text"},
		},
		"elements": []map[string]interface{}{
			{
				"tag": "div", "text": map[string]interface{}{
					"content": fmt.Sprintf(
						"标题: **%s**\n优先级: **%s**\n申请人: **%s**\n申请时间: **%s**\n最近处理时间: **%s**",
						message.Title,
						message.Priority,
						message.Creator,
						message.CreatedAt,
						message.UpdatedAt,
					),
					"tag": "lark_md",
				},
			},
		},
	}
}
