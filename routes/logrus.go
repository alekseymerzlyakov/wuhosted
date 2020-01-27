package routes

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

func Logrus() {
	var filename string = "logfile.log"
	//var filename string = "logfile.log"
	// Create the log file if doesn't exist. And append to it if it already exists.
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	Formatter := new(log.JSONFormatter)
	//f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	//Formatter := new(log.JSONFormatter)
	// You can change the Timestamp format. But you have to use the same date and time.
	// "2006-02-02 15:04:06" Works. If you change any digit, it won't work
	// ie "Mon Jan 2 15:04:05 MST 2006" is the reference time. You can't change it
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	//Formatter.TimestampFormat = "02-01-2006 15:04:05"
	//Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
	if err != nil {
		//log.SetFormatter(Formatter)
		//if err != nil {
		// Cannot open log file. Logging to stderr
		fmt.Println(err)
	} else {
		log.SetOutput(f)
		//	fmt.Println(err)
		//} else {
		//log.SetOutput(f)
	}

	//log.Info("Some info. Earth is not flat" )
	//log.Warning("This is a warning")
	//log.Error("Not fatal. An error. Won't stop execution")
	//log.Fatal("MAYDAY MAYDAY MAYDAY")
	//log.Panic("Do not panic")

}

//}
