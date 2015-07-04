package mail

import (
	"github.com/mailgun/mailgun-go"
	"os"
)

type mailgunProvider struct {
	mg mailgun.Mailgun
}

func MailgunMailer() EmailProvider {

	return &mailgunProvider{
		mg: mailgun.NewMailgun(os.Getenv("MAILGUN_DOMAIN"), os.Getenv("MAILGUN_API_KEY"), os.Getenv("MAILGUN_PUBLIC_API_KEY")),
	}

}

func (p *mailgunProvider) Send(e Email) error {

	msg := mailgun.NewMessage(e.From, e.Subject, e.Content, e.To)

	if len(e.ContentHTML) > 0 {
		msg.SetHtml(e.ContentHTML)
	}

	_, _, err := p.mg.Send(msg)
	return err
}
