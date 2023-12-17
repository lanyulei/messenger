package notification

import (
	"testing"

	"github.com/lanyulei/messenger/sender/dingtalk"

	"github.com/lanyulei/messenger/config"

	"github.com/lanyulei/messenger/sender/dingtalk/common"
	"github.com/stretchr/testify/assert"
)

/*
  @Author : lanyulei
  @Desc :
*/

func TestSend(t *testing.T) {
	var (
		err error
		res dingtalk.UserIdResponse
	)

	// This address is your own profile address
	err = config.FromFile("/Users/mac/lanyulei/project/golang/messenger/config/settings.dev.json")
	assert.Nil(t, err)

	// Get user ID by mobile phone number
	res, err = common.GetDingtalkUserId("188xxxxxxxx")
	assert.Nil(t, err)

	// Send notification
	msg := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": "test",
		},
	}

	_, err = Send(res.Result.UserID, "", false, msg)
	assert.Nil(t, err)
}
