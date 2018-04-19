package store

import (
	"database/sql"
	"log"
)

type Person struct {
	ID   int
	Name string
}

func GetPeople() {
	rows, err := db.Query("select name from people")
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			log.Panic(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Panic(err)
	}
}

func GetOrCreatePerson(name string) (*Person, error) {
	var person Person
	err := db.QueryRow("SELECT id, name FROM people WHERE name=$1", name).Scan(&person.ID, &person.Name)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	} else if err == sql.ErrNoRows {
		return CreatePerson(name)
	}
	return &person, nil
}

func CreatePerson(name string) (*Person, error) {
	var person Person
	_, err := db.Exec("INSERT INTO people (name) VALUES ($1)", name)
	if err != nil {
		log.Panic(err)
	}
	err = db.QueryRow("SELECT id, name FROM people WHERE name=$1", name).Scan(&person.ID, &person.Name)
	if err != nil {
		return nil, err
	}
	return &person, nil
}
