package bridge

// 发送消息的接口
type IMsgSender interface {
	Send(msg string) error
}

// 发送邮件
type EmailMsgSender struct {
	emails []string
}

func NewEmailMsgSender(emails []string) *EmailMsgSender {
	return &EmailMsgSender{emails: emails}
}

func (s *EmailMsgSender) Send(msg string) error {
	return nil
}

// 发送通知的接口
type INotifiction interface {
	Notify(msg string) error
}

// 发送错误的接口
type ErrorNotification struct {
	sender IMsgSender
}

func NewErrorNotification(sender IMsgSender) *ErrorNotification {
	return &ErrorNotification{sender: sender}
}

func (n *ErrorNotification) Notify(msg string) error {
	return n.sender.Send(msg)
}
