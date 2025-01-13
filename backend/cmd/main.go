package main

import (
	"log"
	"net/http"

	"github.com/corsairier/Cat-GPT/backend/pkg/handler"
	"github.com/corsairier/Cat-GPT/backend/pkg/middleware"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Use(middleware.ChatLogger)
	r.HandleFunc("/", handler.Chat)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("The server broke", err)
	}
}
