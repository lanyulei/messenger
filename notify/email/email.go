package email

import (
    "fmt"
    "gopkg.in/gomail.v2"
    "messenger/config"
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
func Send(to, cc []string, title, content string) (err error) {
    var (
        emailConfig = config.GetConfig().Email
    )

    m := gomail.NewMessage()

    // sender
    m.SetHeader("From", emailConfig.User, emailConfig.Alias)
    // receiver
    m.SetHeader("To", to...)
    // copied to
    m.SetHeader("Cc", cc...)
    // title
    m.SetHeader("Subject", title)
    // content
    m.SetBody("text/html", content)
    // attachment
    //m.Attach("./myIpPic.png")

    d := gomail.NewDialer(
        emailConfig.Host,
        emailConfig.Port,
        emailConfig.User,
        emailConfig.Password,
    )

    // send email
    if err = d.DialAndSend(m); err != nil {
        err = fmt.Errorf("mail delivery failure, %s", err.Error())
        return
    }
    return
}
