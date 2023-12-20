package webhook

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/guonaihong/gout"
)

/*
  @Author : lanyulei
  @Desc :
*/

// Send
// @Description: send webhook message
// @param url webhook address
// @param msg message content
// @return err
func Send(urlList []string, msg map[string]interface{}) (err error) {
	if len(urlList) > 0 {
		var (
			wg      sync.WaitGroup
			errList []string
		)

		for _, url := range urlList {
			wg.Add(1)
			go func(url string, msg map[string]interface{}, errList *[]string, wg *sync.WaitGroup) {
				defer wg.Done()

				var response map[string]interface{}
				err = gout.POST(url).
					SetHeader(gout.H{
						"Content-Type": "application/json",
					}).
					SetJSON(msg).
					BindJSON(&response).
					Do()
				if err != nil {
					*errList = append(*errList, fmt.Sprintf("send webhook failed, url: %s, err: %s", url, err.Error()))
					return
				}

				if d, ok := response["errcode"]; ok {
					if int(d.(float64)) != 0 {
						*errList = append(*errList, fmt.Sprintf("send webhook failed, url: %s, errcode: %d, errmsg: %s", url, int(d.(float64)), response["errmsg"].(string)))
						return
					}
				}
			}(url, msg, &errList, &wg)
		}
		wg.Wait()

		if len(errList) > 0 {
			err = errors.New(strings.Join(errList, ", "))
		}
	}
	return
}
