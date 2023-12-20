package dingtalk

import (
	"errors"
	"testing"

	"github.com/lanyulei/messenger/config"

	"github.com/stretchr/testify/assert"
)

/*
  @Author : lanyulei
  @Desc :
*/

func TestGetAccountToken(t *testing.T) {
	var (
		err error
		at  string
	)

	// This address is your own profile address
	err = config.FromFile("/Users/mac/lanyulei/project/golang/messenger/config/settings.dev.json")
	assert.Nil(t, err)

	at, err = GetAccountToken()
	assert.Nil(t, err)

	if at == "" {
		assert.Nil(t, errors.New("get account token failed"))
	}
}

func TestGetDingtalkUserId(t *testing.T) {
	var (
		err error
	)

	// This address is your own profile address
	err = config.FromFile("/Users/mac/lanyulei/project/golang/messenger/config/settings.dev.json")
	assert.Nil(t, err)

	_, err = GetDingtalkUserId("188xxxxxxxx")
	assert.Nil(t, err)
}
