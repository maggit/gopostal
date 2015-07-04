package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jrallison/go-workers"
	"os"
)

func main() {
	redisNamespace := os.Getenv("REDIS_NAMESPACE")
	redisServer := os.Getenv("REDIS_SERVER")
	redisDB := os.Getenv("REDIS_DB")
	redisPool := os.Getenv("REDIS_POOL")
	redisPwd := os.Getenv("REDIS_PWD")

	if (redisServer == "") || (redisDB == "") || (redisPool == "") || (redisNamespace == "") {
		log.Error("Please Start Worker with Required Arguments")
		return
	}

	workers.Configure(map[string]string{
		// location of redis instance
		"namespace": redisNamespace,
		"server":    redisServer,
		// instance of the database
		"database": redisDB,
		// redis password
		"password": redisPwd,
		// number of connections to keep open with redis
		"pool": redisPool,
		// unique process id for this instance of workers (for proper recovery of inprogress jobs on crash)
		"process": "1",
	})

	// pull messages from "myqueue" with concurrency of 10
	workers.Process("sendemail", SendEmail, 10)
	//workers.Process("sendbatch", SendBatch, 10)

	// stats will be available at http://localhost:8080/stats
	go workers.StatsServer(5000)

	// Blocks until process is told to exit via unix signal
	workers.Run()
}
