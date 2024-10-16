package database

import (
	"fmt"
	"log"
	"rapidart/internal/models"
)

func AddLikeToPost(newLike models.Like) error {
	sqlInsert := `
		INSERT INTO rapidart.Like (
		                  UserId,
		                  PostId
		) VALUES (?, ?);`

	_, err := db.Exec(sqlInsert,
		newLike.UserId,
		newLike.PostId,
	)
	if err != nil {
		log.Println("Error: ", err)
		fmt.Println(err)
		return fmt.Errorf("ERROR: %v", err)
	}
	return nil
}
