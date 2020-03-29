package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type tweeterManager struct {
	client *twitter.Client
	creds  credentialmanager
}

func initialize() {
	// 1. Get credentials

	// 2. register client
}

func (t *tweeterManager) readCredentailsFromJSON() {
	// Read file
	file, _ := ioutil.ReadFile("credentials.json")
	err := json.Unmarshal(file, &t.creds)
	if err != nil {
		panic(err)
	}
	fmt.Println("readCredentailsFromJSON")
	fmt.Println(t)
}

func (t *tweeterManager) register() {
	fmt.Println("register")
	fmt.Println(t)
	config := oauth1.NewConfig(t.creds.ConsumerKey, t.creds.ConsumerSecret)
	token := oauth1.NewToken(t.creds.AccessToken, t.creds.AccessSecret)

	/*
		config := oauth1.NewConfig("4HftecNl5ehQqSMIpmblprP5w", "TzUZqnKSpi1WfmO3K02hS0A6P0X7V4Er2qshOpbM8a4jMcUZ40")
		token := oauth1.NewToken("43485947-u0R8lNcvOOvfsR8WsVVSNlBreS7dP7cgnRUgvKgqV", "ZZmkGIBgH8Dn1C02VuXmqAIJ9xALWhVMt6qdFhoL1EZ3g")
		httpClient := config.Client(oauth1.NoContext, token)
	*/

	httpClient := config.Client(oauth1.NoContext, token)

	// config := &clientcredentials.Config{
	// 	ClientID:     t.creds.AccessToken,
	// 	ClientSecret: t.creds.AccessSecret,
	// 	TokenURL:     "https://api.twitter.com/oauth2/token",
	// }
	// // http.Client will automatically authorize Requests
	// httpClient := config.Client(oauth2.NoContext)

	// Twitter client
	t.client = twitter.NewClient(httpClient)
	t.test()
}

func (t *tweeterManager) test() {
	// Home Timeline
	// Send a Tweet
	tweet, resp, err := t.client.Statuses.Update("just setting up my twttr, not really", nil)

	fmt.Println(tweet, resp, err)
}
