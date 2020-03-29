package main

import (
	"fmt"

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

func (t tweeterManager) getcredentials() {
	t.creds = readCredentailsFromJSON()
}

func (t tweeterManager) register() {
	config := oauth1.NewConfig(t.creds.ConsumerKey, t.creds.ConsumerSecret)
	token := oauth1.NewToken(t.creds.AccessToken, t.creds.AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	t.client = twitter.NewClient(httpClient)
	t.test()
}

func (t tweeterManager) test() {
	// Home Timeline
	tweets, resp, err := t.client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: 20,
	})

	fmt.Println(tweets, resp, err)
}
