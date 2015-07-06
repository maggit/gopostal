package mail

import (
	"github.com/sendgrid/sendgrid-go"
	"os"
)

type sendgridProvider struct {
	message string
}

func SendgridMailer() EmailProvider {

	return &sendgridProvider{
		message: "start sendgrid",
	}
}

func (m *sendgridProvider) Send(e Email) error {
	sg := sendgrid.NewSendGridClient(os.Getenv("SENDGRID_USER"), os.Getenv("SENDGRID_KEY"))
	msg := sendgrid.NewMail()
	msg.AddTos(e.To)
	msg.SetSubject(e.Subject)
	msg.SetHTML(e.ContentHTML)
	msg.SetFrom(e.From)
	err := sg.Send(msg)
	return err
}
