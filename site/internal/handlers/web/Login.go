package web

import (
	"log"
	"net/http"
)

type UserInfo struct {
	Name string
	Age  int
}

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		loginGetRequest(w, r)
	default:
		http.Error(w, "this method is not supported", http.StatusNotImplemented)
	}
}

func loginGetRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello LoginGetRequest")

}
