package service

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecretKey = []byte("very-secret-key")

func CreateToken(nameUser string, w http.ResponseWriter, r *http.Request) {
	payload := jwt.MapClaims{
		"sub": nameUser,
		"exp": time.Now().Add(time.Minute * 10).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	hashedToken, _ := token.SignedString(jwtSecretKey)
	fmt.Println(hashedToken)
	livingTime := 1 * time.Minute
	expiration := time.Now().Add(livingTime)

	cookie := http.Cookie{Name: "token", Value: url.QueryEscape(hashedToken), Expires: expiration}
	http.SetCookie(w, &cookie)

}
