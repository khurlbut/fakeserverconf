package fakeserverconf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Path   string
	Status int
	Body   string
}

type Configuration []*Config

func DefaultConfig() Configuration {
	c := &Config{
		Path:   "/",
		Status: 200,
		Body:   "Welcome to fakeserver. This is the Default configuration for Root path (/)",
	}
	configs := []*Config{c}

	return configs
}

func ReadJSON(json string) Configuration {
	return unmarshal([]byte(json))
}

func ReadJSONFile(path string) Configuration {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return unmarshal(bytes)
}

func unmarshal(bytes []byte) Configuration {
	var c Configuration
	err := json.Unmarshal(bytes, &c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}
