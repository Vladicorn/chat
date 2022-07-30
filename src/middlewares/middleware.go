package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const secretKey = "secret"

type ClaimsWithScope struct {
	jwt.StandardClaims
	Scope string
}

func LoggingMiddlewareAuth(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("jwt")
		if err != nil {
			fmt.Println(err)
		} else {
			token, err := jwt.ParseWithClaims(cookie.Value, &ClaimsWithScope{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			})
			if err != nil || !token.Valid {
				http.Error(w, "Forbidden", http.StatusForbidden)
			}
			payload := token.Claims.(*ClaimsWithScope)

			log.Println(payload)
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
		}
	})
}

func GenerateJWC(id uint) (string, error) {
	payload := ClaimsWithScope{}

	payload.Subject = strconv.Itoa(int(id))
	payload.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
	return jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(secretKey))
}
