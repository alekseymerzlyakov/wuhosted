package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func Subscription(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Hello from a Alex!\n")

	buf, bodyErr := ioutil.ReadAll(r.Body) //вычитываем запрос
	if bodyErr != nil {
		log.Print("bodyErr ->      ", bodyErr.Error())
		http.Error(w, bodyErr.Error(), http.StatusNoContent)
		return
	}
	rdr2 := bytes.NewBuffer(buf)
	log.Println("func IndexApproveSend r.URL ---->>>>>>       ", r.URL)
	log.Println("func IndexApproveSend r.Header ---->>>>>>       ", r.Header)
	log.Println("func IndexApproveSend: Request Body ---->>>>>>       ", string(buf))
	fmt.Println("POST - > ", string(buf))
	log.Debug(rdr2)

}

func SubscriptionGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Hello from a Alex!\n Method GET")

	buf, bodyErr := ioutil.ReadAll(r.Body) //вычитываем запрос
	if bodyErr != nil {
		log.Print("bodyErr ", bodyErr.Error())
		http.Error(w, bodyErr.Error(), http.StatusOK)
		return
	}
	rdr2 := bytes.NewBuffer(buf)
	log.Info("GET - > ", rdr2)
	fmt.Println("GET - > ", rdr2)

}

func Time() string {
	now := time.Now() // current local time
	X_Time_ := now.Unix()
	X_Time := strconv.FormatInt(X_Time_, 10)
	fmt.Println("X_Time      " + X_Time)
	return X_Time
}

func sha_256(combined string) string {
	h := sha256.New()
	h.Write([]byte(combined))
	return hex.EncodeToString(h.Sum(nil))
}
