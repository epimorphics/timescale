package auth

import (
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/epimorphics/timescale/backend/store"
	"github.com/gorilla/context"
	"log"
	"net/http"
	"os"
	"strings"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JwtToken struct {
	Token string `json:"token"`
}

func CreateTokenEndpoint(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Panic(err)
		return
	}
	if !store.ValidateUser(user.Username, user.Password) {
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
	})
	tokenString, error := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if error != nil {
		log.Println(error)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
}

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, errors.New("There was an error")
					}
					return []byte(os.Getenv("JWT_SECRET")), nil
				})
				if error != nil {
					w.WriteHeader(http.StatusUnauthorized)
					json.NewEncoder(w).Encode(error.Error())
					return
				}
				if token.Valid {
					context.Set(r, "decoded", token.Claims)
					next(w, r)
				} else {
					w.WriteHeader(http.StatusUnauthorized)
					json.NewEncoder(w).Encode("Invalid authorization token")
				}
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("An authorization header is required")
		}
	})
}
