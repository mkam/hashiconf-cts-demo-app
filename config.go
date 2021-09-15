package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port   int    `json:port`
	Name   string `json:"name"`
	APIURL string `json:"api_url"`
}

func readConfig(filename string) (*Config, error) {
	c := Config{}
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		return nil, err
	}
	return &c, err
}
