package webhook

import (
	"fmt"
	"github.com/guonaihong/gout"
)

/*
  @Author : lanyulei
  @Desc :
*/

// Send
// @Description: send dingtalk robot message
// @param url webhook address
// @param msg message content
// @return err
func Send(urlList []string, msg map[string]interface{}) (err error) {
	if len(urlList) > 0 {
		go func(urlList []string, msg map[string]interface{}) {
			for _, url := range urlList {
				err = gout.POST(url).
					SetHeader(gout.H{
						"Content-Type": "application/json",
					}).
					SetJSON(msg).
					Do()
				if err != nil {
					err = fmt.Errorf("send webhook failed, url: %s, err: %s", url, err.Error())
				}
			}
		}(urlList, msg)
	}
	return
}
