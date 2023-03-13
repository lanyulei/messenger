package webhook

import (
    "github.com/guonaihong/gout"
)

/*
  @Author : lanyulei
  @Desc :
*/

// Send
// @Description: send wecom robot message
// @param url webhook address
// @param msg message content
// @return err
func Send(url string, msg map[string]interface{}) (err error) {
    // request webhook
    err = gout.POST(url).
            SetHeader(gout.H{
                "Content-Type": "application/json",
            }).
        SetJSON(msg).
        Do()
    return
}
