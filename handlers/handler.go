package handlers

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "注册成功")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "登陆成功！")
}
