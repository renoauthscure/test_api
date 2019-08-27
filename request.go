package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func DoRequest(urlTarget string, jsonStr []byte) (string, error) {

	req, err := http.NewRequest("POST", urlTarget, bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, errReq := client.Do(req)
	if err != errReq {
		return "", errReq
	}

	defer resp.Body.Close()

	body, errParse := ioutil.ReadAll(resp.Body)
	if errParse != nil {
		return "", errParse
	}

	return string(body), nil

}
