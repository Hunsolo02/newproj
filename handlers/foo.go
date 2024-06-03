package handlers

import (
	"net/http"
)

func HandleFoo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("foo"))

}
