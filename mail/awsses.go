package mail

import (
	log "github.com/Sirupsen/logrus"
	"github.com/sourcegraph/go-ses"
	"os"
)

type awssesConfig struct {
	amazon ses.Config
}

func AwsSesMailer() EmailProvider {
	// EnvConfig takes the access key ID and secret access key values from the environment variables
	// $AWS_ACCESS_KEY_ID and $AWS_SECRET_KEY, respectively.
	return &awssesConfig{
		amazon: ses.Config{
			Endpoint:        os.Getenv("AWS_END_POINT"),
			AccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
			SecretAccessKey: os.Getenv("AWS_SECRET_KEY"),
		},
	}
}

func (p *awssesConfig) Send(e Email) error {
	AMZConfig := p.amazon
	// This doesn't allow to send email to an array of recipients
	_, err := AMZConfig.SendEmailHTML(e.From, e.To[0], e.Subject, e.Content, e.ContentHTML)
	if err != nil {
		log.Error("Error sending email: %s\n", err)
	}
	return err
}
