package email

import (
	"testing"

	"github.com/lanyulei/messenger/types"

	"github.com/lanyulei/messenger/config"
	"github.com/stretchr/testify/assert"
)

/*
  @Author : lanyulei
  @Desc :
*/

func TestSend(t *testing.T) {
	var (
		err     error
		to      []string
		cc      []string
		title   string
		content *types.Message
	)

	// This address is your own profile address
	err = config.FromFile("/Users/mac/lanyulei/project/golang/messenger/config/settings.dev.json")
	assert.Nil(t, err)

	to = []string{"xxxx@163.com"}

	title = "lanyule-test"
	content = &types.Message{
		Title:     "兰玉磊-测试",
		Priority:  "紧急",
		Creator:   "lanyulei",
		CreatedAt: "2023-03-03 11:40:13",
		UpdatedAt: "2023-03-03 11:40:13",
	}

	err = Send(to, cc, title, content)
	assert.Nil(t, err)
}
