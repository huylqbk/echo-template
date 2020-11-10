package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	MethodGet    = "GET"
	MethodPost   = "POST"
	MethodPut    = "PUT"
	MethodDelete = "DELETE"
	MethodPatch  = "PATCH"
)

func HTTPClient(url string, method string, body map[string]interface{}, token string) ([]byte, int) {
	requestBody, err := json.Marshal(body)
	if err != nil {
		log.Println("Marshal", err)
		return nil, -1
	}
	client := &http.Client{}
	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println("NewRequest", err)
		return nil, -1
	}
	request.Header.Set("Content-Type", "application/json")
	if token != "" {
		request.Header.Set("authorization", token)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Println("ClientDo", err)
		return nil, -1
	}
	code := response.StatusCode
	defer response.Body.Close()
	bodyResp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("ReadBody", err)
		return nil, code
	}
	return bodyResp, code
}
