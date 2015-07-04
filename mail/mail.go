package mail

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

var provider EmailProvider

func init() {
	switch os.Getenv("EMAIL_PROVIDER") {
	case "mailgun":
		provider = MaigunMailer()
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
	if provider == nil {
		return log.Error("No email provider")
	}
	return provider.Send(m)
}
