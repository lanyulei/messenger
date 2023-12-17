package webhook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
  @Author : lanyulei
  @Desc :
*/

func TestSend(t *testing.T) {

	urlList := []string{"https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxx"}
	msg := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": "lanyulei-test",
		},
	}

	err := Send(urlList, msg)
	assert.Nil(t, err)
}
