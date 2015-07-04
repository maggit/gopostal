package mail

import (
	"github.com/sourcegraph/go-ses"
	"os"
)

type awssesProvider struct {
	mg awsses.Mailgun
}

func AwsSesMailer() EmailProvider {
	// EnvConfig takes the access key ID and secret access key values from the environment variables
	// $AWS_ACCESS_KEY_ID and $AWS_SECRET_KEY, respectively.
	return &awssesProvider{
		amazon: ses.Config{
			Endpoint:        os.Getenv("AWS_END_POINT"),
			AccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
			SecretAccessKey: os.Getenv("AWS_SECRET_KEY"),
		},
	}
}

func (p *awssesProvider) Send(e Email) error {
	_, err := amazon.SendEmail(e.From, e.To, e.Subject, e.Content, e.ContentHTML)
	if err != nil {
		log.Error("Error sending email: %s\n", err)
	}
	return err
}
