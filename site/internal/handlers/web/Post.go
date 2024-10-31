package web

import (
	"log"
	"net/http"
	"rapidart/internal/database"
	"rapidart/internal/util"
	"strconv"
)

type Title struct {
	Title  string
	Image  []byte
	PostId int
}

type postTemplateModel struct {
	Image []byte
}

func Post(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		postGetRequest(w, r)
	default:
		http.Error(w, "this method is not supported", http.StatusNotImplemented)
	}
}

func postGetRequest(w http.ResponseWriter, r *http.Request) {

	postIdURL := r.URL.Query().Get("post_id")
	log.Println(postIdURL)
	postId, _ := strconv.Atoi(postIdURL)

	post, err := database.GetPostById(postId)

	headerTitle := Title{
		Title:  "Post",
		Image:  post.Image,
		PostId: post.PostId,
	}

	err = util.HttpServeTemplate("post.tmpl", headerTitle, w)
	if err != nil {
		log.Println(err)
		util.HttpReturnError(http.StatusInternalServerError, w)
		return
	}
}
