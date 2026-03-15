package settings

import (
	"log"
	"os"
	"strconv"
	"time"
)

func Init() {
	var splits int64
	var timeout int64
	var err error
	splits, err = strconv.ParseInt(os.Getenv("APP_WORKERS"), 10, 64)
	if err != nil {
		log.Panicln(err)
	}

	Splits = int(splits)

	timeout, err = strconv.ParseInt(os.Getenv("APP_TIMEOUT_MS"), 10, 64)
	if err != nil {
		log.Panicln(err)
	}

	Splits = int(splits)
	Timeout = time.Duration(timeout) * time.Millisecond

}

var Splits int
var Timeout time.Duration
