package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jrallison/go-workers"
	"github.com/mailgun/mailgun-go"
	"io/ioutil"
	"net/http"
	"os"
)

func SendEmail(msg *workers.Msg) {
	params, err := msg.Args().Array()
	if err != nil {
		log.Error("param is not a string:", err)
		os.Exit(1)
	}

	from := params[0].(string)
	to := params[1].(string)
	subject := params[2].(string)
	plain := params[3].(string)
	emailHtmlUrl := params[4].(string)

	response, err := http.Get(emailHtmlUrl)
	if err != nil {
		log.Error("%s", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error("%s", err)
		os.Exit(1)
	}

	log.Info("Got email body")

	mg := mailgun.NewMailgun(os.Getenv("MAILGUN_DOMAIN"), os.Getenv("MAILGUN_API_KEY"), "")
	m := mg.NewMessage(from, subject, plain)
	m.SetHtml(string(contents))
	m.AddRecipient(to)
	//m.SetDeliveryTime(time.Now().Add(24 * time.Hour))
	m.SetTracking(true)

	_, id, err := mg.Send(m)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	log.Info("Message id=%s", id)

	return
}
