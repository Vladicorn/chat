package handlers

import (
	"chat/src/middlewares"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data User
	_ = json.NewDecoder(r.Body).Decode(&data)
	if data.Login == "admin" && data.Password == "1234" {
		fmt.Println("ok")
		token, err := middlewares.GenerateJWC(1)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "jwt", Value: token, Expires: expiration}
		http.SetCookie(w, &cookie)

	}
}
