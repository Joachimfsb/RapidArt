package web

import (
	"log"
	"net/http"
	"rapidart/internal/util"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		profileGetRequest(w, r)
	default:
		http.Error(w, "this method is not supported", http.StatusNotImplemented)
	}
}

func profileGetRequest(w http.ResponseWriter, r *http.Request) {
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
