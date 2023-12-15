package types

/*
  @Author : lanyulei
  @Desc :
*/

type Message struct {
	Title     string
	Priority  string
	Creator   string
	CreatedAt string
	UpdatedAt string
}

const (
	MarkdownMessageType = "markdown"
	TextMessageType     = "text"
)