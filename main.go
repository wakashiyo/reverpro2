package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	director := func(request *http.Request) {
		//URL of request
		url := *request.URL
		//protocol of request
		url.Scheme = "http"
		//to server
		url.Host = "api:9000"

		//read request body as buffer
		buffer, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Fatal(err.Error())
		}

		//create new request
		req, err := http.NewRequest(request.Method, url.String(), bytes.NewBuffer(buffer))
		if err != nil {
			log.Fatal(err.Error())
		}

		//add reqeust header from originally
		req.Header = request.Header
		*request = *req
	}

	rp := &httputil.ReverseProxy{Director: director}
	server := http.Server{
		Addr:    ":8484",
		Handler: rp,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
