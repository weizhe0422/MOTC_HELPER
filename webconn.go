package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type webconn struct {
	url string
}

func NewWebconn(url string) *webconn {
	wc := new(webconn)
	wc.url = url
	return wc
}

func GetHTTPResponse(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return body, nil
}
