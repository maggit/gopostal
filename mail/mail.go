package mail

import (
	"os"
)

var provider EmailProvider

func init() {
	switch os.Getenv("EMAIL_PROVIDER") {
	case "mailgun":
		provider = MailgunMailer()
	case "sendgrid":
		provider = SendgridMailer()
	case "AWS_SES":
		provider = AwsSesMailer()
	}
}

type Email struct {
	From        string
	Subject     string
	Content     string
	ContentHTML string
	To          []string
}

type EmailProvider interface {
	Send(Email) error
}

func Send(m Email) error {
	return provider.Send(m)
}
