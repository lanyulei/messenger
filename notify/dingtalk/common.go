package dingtalk

import (
    "encoding/json"
    "fmt"
    "github.com/guonaihong/gout"
    "messenger/config"
    "time"
)

/*
  @Author : lanyulei
  @Desc :
*/

// GetAccountToken
// @Description: get dingtalk account token
// @return err
func GetAccountToken() (at string, err error) {
    var (
        result         []byte
        dingTalkConfig = config.GetConfig().DingTalk
    )

    if accessToken == nil || time.Now().Unix() > accessToken["expires_time"].(int64) {
        err = gout.GET(AccessTokenURL).SetQuery(gout.H{
            "appkey":    dingTalkConfig.AppKey,
            "appsecret": dingTalkConfig.AppSecret,
        }).BindBody(&result).Do()
        if err != nil {
            err = fmt.Errorf("failed to get access token, err:%s\n", err.Error())
            return
        }

        err = json.Unmarshal(result, &accessToken)
        if err != nil {
            err = fmt.Errorf("failed to get access token, err:%s\n", err.Error())
            return
        }

        if errCode, ok := accessToken["errcode"]; ok && int(errCode.(float64)) != 0 {
            err = fmt.Errorf("failed to get dingtalk access token, err:%s\n", accessToken["errmsg"].(string))
            return
        }

        accessToken["expires_time"] = time.Now().Unix() + int64(accessToken["expires_in"].(float64))
    }

    at = accessToken["access_token"].(string)

    return
}
