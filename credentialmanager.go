package main

import (
	"encoding/json"
	"io/ioutil"
)

type credentialmanager struct {
	ConsumerKey    string `json:"consumerKey"`
	ConsumerSecret string `json:"consumerSecret"`
	AccessToken    string `json:"accessToken"`
	AccessSecret   string `json:"accessSecret"`
}

func readCredentailsFromJSON() credentialmanager {
	// Read file
	var creds credentialmanager
	file, _ := ioutil.ReadFile("credentials.json")
	err := json.Unmarshal(file, &creds)
	if err != nil {
		panic(err)
	}
	return creds
}
