package postPic

import (
	"log"
	"rapidart/internal/database"
)

// Fetches a user's profile picture by their ID
func GetPostPic(postId int) ([]byte, error) {
	post, err := database.GetPostById(postId)
	if err != nil {
		log.Println("Error fetching profile picture:", err)
		return nil, err
	}

	return post.Image, nil
}
