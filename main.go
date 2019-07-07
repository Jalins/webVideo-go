package main

import (
	"log"
	"net/http"
	"webVideo-go/handlers"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	r.POST("/user", handlers.CreateUser)
	r.POST("/login", handlers.Login)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Panic(err)
	}
}
