package fakeserverconf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Configuration site configuration structure
type Configuration struct {
	IPAddress string
	Port      string
	Pages     []*Page
}

// Page page configuration structure
type Page struct {
	Path    string
	Status  int
	Body    string
	Headers []string
	Cookies []string
}

// DefaultConfig configuruation to use if none other is specified
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

// ReadJSON read configuration from JSON string
func ReadJSON(json string) Configuration {
	return unmarshal([]byte(json))
}

// ReadJSONFile read configuration from JSON file
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
