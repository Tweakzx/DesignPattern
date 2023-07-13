package template

import "fmt"

// 发送短信的模板
type ISMS interface {
	send(content string, phone int) error
}

type sms struct {
	ISMS
}

func (s *sms) Validate(content string) error {
	if len(content) > 63 {
		return fmt.Errorf("content is too long")
	}
	return nil
}

func (s *sms) Send(content string, phone int) error {
	if err := s.Validate(content); err != nil {
		return err
	}
	return s.send(content, phone)
}

// 中国联通短信
type UnicomSms struct {
	*sms //继承sms的Validate和Send方法
}

// 补全模板中需要的send方法
func (u *UnicomSms) send(content string, phone int) error {
	fmt.Printf("unicom send sms: %s, phone: %d\n", content, phone)
	return nil
}

func NewUnicomSms() *UnicomSms {
	uni := &UnicomSms{}
	uni.sms = &sms{ISMS: uni}
	return uni
}
