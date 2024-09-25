package service

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/golang-jwt/jwt"
)

func ReadCookie(name string, r *http.Request) (value string, err error) {
	if name == "" {
		return value, errors.New("you are trying to read empty cookie")
	}
	cookie, err := r.Cookie(name)
	if err != nil {
		return value, err
	}
	str := cookie.Value
	value, _ = url.QueryUnescape(str)
	return value, err
}

func CheckToken(r *http.Request) (bool, string) {
	tokenS, err := ReadCookie("token", r)
	if err != nil {
		fmt.Println("Токена нет")
		return false, ""
	}

	token, err := jwt.Parse(tokenS, func(t *jwt.Token) (interface{}, error) {
		return []byte("very-secret-key"), nil
	})
	if err != nil || !token.Valid {
		log.Printf("Error parsing token: %v", err)

	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Printf("Error getting claims from token")

	}

	id, ok := claims["sub"].(string)
	if !ok {
		log.Printf("Roles not found in token")

	}

	return true, id
}
