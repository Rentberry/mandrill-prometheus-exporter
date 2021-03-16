package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func getMandrillUserInfo(apiKey string) (*mandrillUserInfo, error) {

	body := bytes.Buffer{}
	body.WriteString("{\"key\": \"")
	body.WriteString(apiKey)
	body.WriteString("\"}")

	req, err := http.NewRequest("POST", "https://mandrillapp.com/api/1.0/users/info.json", &body)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	respBody, _ := ioutil.ReadAll(resp.Body)

	result := mandrillUserInfo{}
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
