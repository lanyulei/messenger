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

	mobiles := []string{"188xxxxxxxx"}

	_, err = Send(mobiles, map[string]interface{}{
		"config": map[string]interface{}{
			"wide_screen_mode": true,
		},
		"header": map[string]interface{}{
			"title": map[string]interface{}{
				"tag":     "plain_text",
				"content": "lanyulei-test",
			},
		},
		"elements": []map[string]interface{}{
			{
				"tag": "div",
				"fields": []map[string]interface{}{
					{
						"is_short": true,
						"text": map[string]interface{}{
							"tag":     "plain_text",
							"content": "lanyulei-test",
						},
					},
				},
			},
		},
	})
	assert.Nil(t, err)
}
