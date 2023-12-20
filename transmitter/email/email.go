package email

import (
	"fmt"

	"github.com/lanyulei/messenger/config"
	"gopkg.in/gomail.v2"
)

/*
  @Author : lanyulei
  @Desc :
*/

// Send
// @Description: send email
// @param to receiver
// @param cc copied to
// @param title email title
// @param content email content
// @return err
func Send(to, cc []string, title string, contentType, content string) (err error) {
	if contentType == "" {
		contentType = "text/html"
	}

	m := gomail.NewMessage()

	// sender
	m.SetHeader("From", config.GetConfig().Email.User)
	// receiver
	m.SetHeader("To", to...)
	// copied to
	m.SetHeader("Cc", cc...)
	// title
	m.SetHeader("Subject", title)
	// content
	m.SetBody(contentType, content)
	// attachment
	//m.Attach("./myIpPic.png")

	d := gomail.NewDialer(
		config.GetConfig().Email.Host,
		config.GetConfig().Email.Port,
		config.GetConfig().Email.User,
		config.GetConfig().Email.Password,
	)

	// send email
	if err = d.DialAndSend(m); err != nil {
		err = fmt.Errorf("mail delivery failure, %s", err.Error())
		return
	}
	return
}
