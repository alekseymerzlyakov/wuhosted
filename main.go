package main

import (
	"fmt"
	"github.com/magiconair/properties"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {

	var filename string = "logfile.log"
	f, _ := os.Create(filename)
	// Create the log file if doesn't exist. And append to it if it already exists.
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	Formatter := new(log.TextFormatter)
	// You can change the Timestamp format. But you have to use the same date and time.
	// "2006-02-02 15:04:06" Works. If you change any digit, it won't work
	// ie "Mon Jan 2 15:04:05 MST 2006" is the reference time. You can't change it
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
	if err != nil {
		// Cannot open log file. Logging to stderr
		fmt.Println(err)
	} else {
		log.SetOutput(f)
	}

	http.HandleFunc("/wu", IndexHandler)
	http.HandleFunc("/approvesend", IndexApproveSend)
	http.HandleFunc("/subscription", Subscription)
	//http.HandleFunc("/subscription2", Subscription)
	//http.HandleFunc("/subscription", Subscription)
	//r := mux.NewRouter()
	//r.HandleFunc("/subscription", Subscription).Methods("POST") //https://proglib.io/p/rest-api-go
	//r.HandleFunc("/subscription", SubscriptionGet).Methods("GET")
	//r.HandleFunc("/subscription", SubscriptionPut).Methods("PUT")

	http.ListenAndServe(":8081", nil)
}

//Open Properties file
func GetString(key string) string {
	p := properties.MustLoadFile(path()+"/wu.properties", properties.UTF8)
	return p.MustGet(key)
}
