package main

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"

	"net/http"
)


func Subscription(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	buf, bodyErr := ioutil.ReadAll(r.Body) //вычитываем запрос
	if bodyErr != nil {
		log.Print("bodyErr ", bodyErr.Error())
		http.Error(w, bodyErr.Error(), http.StatusInternalServerError)
		return
	}

	//rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf)) // конвертируем в string
	rdr2 := bytes.NewBuffer(buf)
	//log.Info("log.Info - > ", rdr1)
	//fmt.Println("fmt.Println - > ", rdr1)

	log.Info("log.Info rdr2 - > ", rdr2)
	fmt.Println("fmt.Println rdr2 - > ", rdr2)


}