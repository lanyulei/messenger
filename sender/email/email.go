package email

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/lanyulei/messenger/config"
	"github.com/lanyulei/messenger/types"
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
func Send(to, cc []string, title string, content *types.Message) (err error) {
	var (
		value string
	)

	value, err = formatContent(content)
	if err != nil {
		err = fmt.Errorf("format content failure, %s", err.Error())
		return
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
	m.SetBody("text/html", value)
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

func formatContent(message *types.Message) (content string, err error) {
	var (
		buf bytes.Buffer
	)

	tmpl, err := template.New("email").Parse(templateData)
	if err != nil {
		return
	}

	err = tmpl.Execute(&buf, message)
	if err != nil {
		return
	}

	content = buf.String()
	return
}
