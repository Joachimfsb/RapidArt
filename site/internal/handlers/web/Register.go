package web

import (
	"log"
	"net/http"
	"rapidart/internal/util"
)

func Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		registerGetRequest(w, r)
	default:
		http.Error(w, "this method is not supported", http.StatusNotImplemented)
	}
}

func registerGetRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello RegisterGetRequest")

	var model = User{
		Name: "Bob",
		Age:  30,
	}

	err := util.HttpServeTemplate("profile.html", model, w)
	if err != nil {
		log.Println(err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}
}
