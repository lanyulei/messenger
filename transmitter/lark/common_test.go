package lark

import (
	"testing"

	"github.com/lanyulei/messenger/config"
	"github.com/stretchr/testify/assert"
)

/*
  @Author : lanyulei
  @Desc :
*/

// GetLarkAccountToken
// @Description: get lark tenant account token
// @return err
func TestGetLarkAccountToken(t *testing.T) {
	err := config.FromFile("/Users/mac/lanyulei/project/golang/messenger/config/settings.dev.json")
	assert.Nil(t, err)

	_, err = GetLarkAccountToken()
	assert.Nil(t, err)
}

func TestGetLarkUserIDByMobiles(t *testing.T) {
	err := config.FromFile("/Users/mac/lanyulei/project/golang/messenger/config/settings.dev.json")
	assert.Nil(t, err)

	_, err = GetLarkUserIDByMobiles([]string{"188xxxxxxxx"})
	assert.Nil(t, err)
}
