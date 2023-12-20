package common

import (
	"fmt"
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
		assert.Nil(t, fmt.Errorf("get account token failed"))
	}
}

func TestGetWeComUserId(t *testing.T) {
	// This address is your own profile address
	err := config.FromFile("/Users/mac/lanyulei/project/golang/messenger/config/settings.dev.json")
	assert.Nil(t, err)

	_, err = GetWeComUserId("188xxxxxxxx")
	assert.Nil(t, err)
}
