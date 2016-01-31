package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type oauthConfig struct {
	ClientID     string
	ClientSecret string
}

type configObject struct {
	Github   oauthConfig
	Google   oauthConfig
	Facebook oauthConfig
}

func loadConfig() *configObject {
	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		log.Fatalln("Failed to read config.json:", e)
	}
	var config configObject
	json.Unmarshal(file, &config)
	return &config
}
