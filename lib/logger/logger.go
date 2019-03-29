package logger

import (
	"flag"
	"log"
	"os"
	"time"
)

// Log saves logs to specific file.
var Log *log.Logger

func init() {

	t := string(time.Now().Format("2006-01-02"))

	// set location of log file
	var logpath = "rabbitmq_" + t + ".log"

	flag.Parse()
	var file, err1 = os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err1 != nil {
		panic(err1)
	}

	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
}
