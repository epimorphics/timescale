package store

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func ValidateUser(username, password string) bool {
	var hashedPassword string
	err := db.QueryRow("select password from users WHERE username=$1", username).Scan(&hashedPassword)
	if err != nil {
		log.Panic(err)
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err == nil {
		return true
	}
	log.Panic(err)
	return false
}

func CreateUser(username string, password string) error {
	hashed, err := hashAndSalt(password)
	if err != nil {
		log.Panic(err)
		return err
	}
	_, err = db.Exec("INSERT INTO users (username, password) VALUES ($1, $2) ON CONFLICT (username) DO UPDATE set password = excluded.password", username, hashed)
	if err != nil {
		log.Panic(err)
		return err
	}
	return nil
}

func hashAndSalt(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
