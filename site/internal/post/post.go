package post

import (
	"rapidart/internal/database"
	"rapidart/internal/models"
	"time"
)

// Create a new post
func CreatePost(userId, basisCanvasId int, image []byte, caption string, timeSpent int) error {

	post := models.Post{
		UserId:           userId,
		BasisCanvasId:    basisCanvasId,
		Image:            image,
		Caption:          caption,
		TimeSpentDrawing: timeSpent,
		CreationDateTime: time.Now(),
		Active:           true,
	}

	err := database.AddPost(post)
	if err != nil {
		return err
	}

	return nil
}
