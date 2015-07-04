package main

import (
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/jrallison/go-workers"
	_ "github.com/lib/pq"
	"os"
	"strconv"
	"time"
)

func SendBatch(message *workers.Msg) {
	return
}
