package utils

import (
	"bytes"
	"log"
	"net/http"
	"time"
)

func HTTPRequest(url string, method string, payload []byte) *http.Response {
	var request *http.Request
	var err error
	var body *bytes.Buffer
	timeout := time.Duration(5 * time.Second)
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{
		Timeout:   timeout,
		Transport: tr,
	}
	switch method {
	case "GET":
		body = bytes.NewBuffer(nil)
	case "POST":
		body = bytes.NewBuffer(payload)
	case "PUT":
		body = bytes.NewBuffer(payload)
	case "DELETE":
		body = bytes.NewBuffer(nil)
	default:
		body = bytes.NewBuffer(nil)
	}
	if err != nil {
		log.Println("Create request error.", err.Error())
	}
	request, err = http.NewRequest(method, url, body)
	if err != nil {
		log.Println("Client request to "+url+" is not success.", err.Error())
	}
	request.Header.Add("Content-type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		log.Println("Client is error.", err.Error())
	}
	return response
}
