package main

import (
	"github.com/epimorphics/timescale/backend/store"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
)

type Specification struct {
	Port               string `envconfig:"PORT" default:"3000"`
	AdminUsername      string `split_words:"true"`
	AdminPassword      string `split_words:"true"`
	JWTSecret          string `required:"true" envconfig:"JWT_SECRET"`
	DatabaseConnection string `envconfig:"DB_CONN" required:"true"`
}

var s Specification

func main() {
	err := envconfig.Usage("timescale", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = envconfig.Process("timescale", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	store.OpenDB(s.DatabaseConnection)
	if s.AdminUsername != "" && s.AdminPassword != "" {
		err := store.CreateUser(s.AdminUsername, s.AdminPassword)
		if err != nil {
			log.Panic(err)
			return
		}
	}
	defer store.CloseDB()
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":"+s.Port, router))
}
