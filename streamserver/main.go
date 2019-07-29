package main

import (
	"net/http"
	"streamserver/handler"
	"streamserver/helper"

	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *helper.ConnLimiter
}

func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = helper.NewConnLimiter(cc)
	return &m
}

func (m *middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		helper.SendErrorResponese(w, http.StatusTooManyRequests, "请求太多")
		return
	}

	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.GET("/videos/:vid-id", handler.StreamHandler)

	router.POST("/upload/:vid-id", handler.UploadHandler)

	return router
}

func main() {
	r := RegisterHandler()
	wareHandler := NewMiddleWareHandler(r, 2)
	http.ListenAndServe(":9000", wareHandler)
}
