package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/mailgun/mailgun-go"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get(os.Getenv("TEST_EMAIL_URL"))
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
	mg := mailgun.NewMailgun(os.Getenv("SANDBOX_MAILGUN_DOMAIN"), os.Getenv("SANDBOX_MAILGUN_API_KEY"), "")
	m := mg.NewMessage(os.Getenv("TEST_FROM"), os.Getenv("TEST_SUBJECT"), os.Getenv("TEST_PLAIN"))
	m.SetHtml(string(contents))
	m.AddRecipient(os.Getenv("TEST_TO"))
	m.SetTracking(true)
	_, id, err := mg.Send(m)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	log.Printf("Message id=%s", id)
}
