package notify

/*
  @Author : lanyulei
  @Desc :
*/

type Interface interface {
    Email(to, cc []string, title, content string) error
}

type notify struct{}

func New() Interface {
    return &notify{}
}

// Email
// @Description: email notification
// @param to receiver
// @param cc copied to
// @param title email title
// @param content email content
// @return err
func (n *notify) Email(to, cc []string, title, content string) (err error) {
    return nil
}
