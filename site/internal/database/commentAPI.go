package database

import (
	"fmt"
	"log"
	"rapidart/internal/models"
)

func AddCommentToPost(newComment models.Comment) error {
	sqlInsert := `
		INSERT INTO Comment (
		                  UserId,
		                  PostId,
		                  Message,
		                  CreationDateTime
		) VALUES (?, ?, ?, ?);`

	_, err := db.Exec(sqlInsert,
		newComment.UserId,
		newComment.PostId,
		newComment.Message,
		newComment.CreationDateTime,
	)
	if err != nil {
		log.Println("Error: ", err)
		fmt.Println(err)
		return fmt.Errorf("ERROR: %v", err)
	}
	return nil
}
