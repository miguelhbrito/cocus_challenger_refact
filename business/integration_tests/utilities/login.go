package utilities

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/cocus_challenger_refact/business/data/login"
)

func CreateUser() {
	login, _ := json.Marshal(login.NewLogin{Username: User, Password: Password})

	req, err := http.NewRequest(http.MethodPost, "http://"+BaseURL+"/login/create", bytes.NewBuffer(login))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return
}

func GenerateToken() {
	loginReq, _ := json.Marshal(login.NewLogin{Username: User, Password: Password})

	req, err := http.NewRequest(http.MethodPost, "http://"+BaseURL+"/login", bytes.NewBuffer(loginReq))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	dec.DisallowUnknownFields()

	var respLogin login.Token
	err = dec.Decode(&respLogin)
	if err != nil {
		panic(err)
	}

	Token = respLogin.Token
	return
}
