package email

import (
	"code.project.com/InstantMessaging/pkg/config"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

// SendEmailCode 发送邮箱验证码
func SendEmailCode(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("Get <%s>", config.Config.Email.Username)
	e.To = []string{toUserEmail}
	e.Subject = "验证码已发送，请查收"
	e.HTML = []byte(fmt.Sprintf("您的验证码:<b> %s </b>", code))
	return e.Send(
		config.Config.Email.Addr,
		smtp.PlainAuth("", config.Config.Email.Username, config.Config.Email.Password, config.Config.Email.Host),
	)
}
