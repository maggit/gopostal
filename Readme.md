# gopostal
**gopostal** is a personal exercise to learn **[golang](https://golang.org/)** and minimize job latency and cost while processing jobs. **(WORK IN PROGRESS)**

![Gopher](https://golang.org/doc/gopher/frontpage.png)



## The problem

We have a ruby application that schedules email delivery jobs using sidekiq, the idea is to process this emails using a go-based background worker.

### Getting Started

1. Install Goop from outside the proejct: `go get github.com/nitrous-io/goop`

2. Run `goop install`

3. Build worker `goop go build worker.go`

4. Run the worker `./worker -queues=hello`


### Libraries Used

* [Go Worker](http://github.com/benmanns/goworker) - This is the main Sidekiq/Resque compatible worker library.

* [MAILGUN](https://github.com/mailgun/mailgun-go) - Go library for sending mail with the Mailgun API.


## Examples
*First, you need the following env variables:*

` MAILGUN_DOMAIN `

` MAILGUN_API_KEY `

### Run the Playground
Test email delivery:

`goop go run playground.go`

*You need the following env variables:*

` SANDBOX_MAILGUN_DOMAIN `

` SANDBOX_MAILGUN_API_KEY  `

` TEST_FROM `

` TEST_SUBJECT `

` TEST_PLAIN `

` TEST_TO `

` TEST_EMAIL_URL `

## TODO
* Add send by batch support (needs to connect to db)
* Complete examples
* Write tests
* Add Sendgrid support
* Add Amazon AWS support (https://github.com/nathanwinther/go-awsses, https://github.com/sourcegraph/go-ses)


## LICENSE
MIT

