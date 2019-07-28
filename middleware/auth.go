package middleware

import (
	"net/http"
	"webVideo-go/helper"
	"webVideo-go/session"
)

var (
	HEADER_FIELD_SESSION = "X-Session-Id"
	HEADER_FIELD_UNAME   = "X-User-Name"
)

func ValidateUserSession(r *http.Request) bool {

	sid := r.Header.Get(HEADER_FIELD_SESSION)

	if len(sid) == 0 {
		return false
	}

	uname, ok := session.IsSeesionExpired(sid)
	if ok {
		return false
	}

	r.Header.Add(HEADER_FIELD_UNAME, uname)
	return true
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FIELD_UNAME)

	if len(uname) == 0 {
		helper.SendErroeResponse(w, helper.ErrorNotAuthUser)
		return false
	}

	return true

}
