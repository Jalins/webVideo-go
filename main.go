package main

import (
	"fmt"
	"log"
	"net/http"
	"webVideo-go/handlers"
	"webVideo-go/middleware"

	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}

	m.r = r
	return &m

}

func (m *middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check session
	middleware.ValidateUserSession(r)
	m.r.ServeHTTP(w, r)
}

func main() {
	r := httprouter.New()

	r.POST("/user", handlers.CreateUser)
	r.POST("/login", handlers.Login)
	wareHandler := NewMiddleWareHandler(r)
	err := http.ListenAndServe(":8080", wareHandler)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("服务器启动成功。。。")
}
