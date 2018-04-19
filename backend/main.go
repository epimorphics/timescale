package main

import (
	"github.com/epimorphics/timescale/backend/store"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
)

type config struct {
	Port string `yaml:"port"`
}

func (c *config) getConf() *config {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Panic(err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Panic(err)
	}
	return c
}

func main() {
	var c config
	c.getConf()
	store.OpenDB()
	defer store.CloseDB()
	router := NewRouter()
	log.Fatal(http.ListenAndServe(c.Port, router))
}
