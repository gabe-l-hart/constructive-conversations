package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Port        string   `json:"port"`
	Identities  []string `json:"identities"`
	DbFilename  string   `json:"db-file"`
	LogFilename string   `json:"log-file"`
}

func ParseConfig(cfgFile string) Config {
	cfg := Config{}
	if content, err := ioutil.ReadFile(cfgFile); nil != err {
		log.Fatalf("Unable to read config file [%s]: "+err.Error(), cfgFile)
	} else {
		if err := json.Unmarshal(content, &cfg); nil != err {
			log.Fatalf("Unable to parse config file [%s]: "+err.Error(), cfgFile)
		}
	}
	return cfg
}
