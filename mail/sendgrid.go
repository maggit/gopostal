package mail

import (
	"github.com/sendgrid/sendgrid-go"
	"os"
)

type sendgridProvider struct {
	mg mailgun.Mailgun
}

func SendgridMailer() EmailProvider {

	return &sendgridProvider{
		sg: sendgrid.NewSendGridClient(os.Getenv("SENDGRID_USER"), os.Getenv("SENDGRID_KEY")),
	}

}

func (p *sendgridProvider) Send(e Email) error {

	msg := sendgrid.NewMail()
	msg.AddTo(e.To)
	msg.SetSubject(e.Subject)
	msg.SetHTML(e.ContentHTML)
	msg.SetFrom(e.From)

	_, err := p.sg.Send(msg)
	return err
}
