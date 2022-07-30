// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"chat/src/database"
	"chat/src/handlers"
	"chat/src/middlewares"
	"chat/src/webs"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var addr = flag.String("addr", ":8080", "http service address")

func ServeHome(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "home.html")
}

func main() {

	flag.Parse()

	hub := webs.NewHub()
	go hub.Run()

	err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/login", handlers.Login)

	r.Handle("/ws", middlewares.LoggingMiddlewareAuth(http.HandlerFunc(hub.ConnectWS)))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://192.168.0.56:8081"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)
	srv := &http.Server{
		Handler:      handler,
		Addr:         "192.168.0.56:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
