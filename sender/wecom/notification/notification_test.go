package notification

import (
	"testing"

	"github.com/lanyulei/messenger/config"

	"github.com/stretchr/testify/assert"
)

/*
  @Author : lanyulei
  @Desc :
*/

func TestSend(t *testing.T) {
	// This address is your own profile address
	err := config.FromFile("/Users/mac/lanyulei/project/golang/messenger/config/settings.dev.json")
	assert.Nil(t, err)

	msg := map[string]interface{}{
		"touser":  "LanYuLei",
		"msgtype": "text",
		"agentid": config.GetConfig().WeCom.AgentId,
		"text": map[string]interface{}{
			"content": "lanyulei-test",
		},
	}

	_, err = Send(msg)
	assert.Nil(t, err)
}
