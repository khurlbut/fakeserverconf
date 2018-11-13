package fakeserverconf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Configuration struct {
	IPAddress string
	Port      int
	Pages     []*Page
}

type Page struct {
	Path   string
	Status int
	Body   string
}

func DefaultConfig() Configuration {
	p := &Page{
		Path:   "/",
		Status: 200,
		Body:   "Welcome to fakeserver. This is the Default configuration for Root path (/)",
	}
	pages := []*Page{p}

	config := &Configuration{
		Pages: pages,
	}

	return *config
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
