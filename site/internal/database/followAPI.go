package database

import (
	"fmt"
	"log"
	"rapidart/internal/models"
)

func NewFollow(follow models.Follow) error {
	sqlInsert := `
		INSERT INTO Follow (
		                  FollowerUserId,
		                  FolloweeUserId
		) VALUES (?, ?);`

	_, err := db.Exec(sqlInsert,
		follow.FollowerUserId,
		follow.FolloweeUserId,
	)
	if err != nil {
		log.Println("Error: ", err)
		fmt.Println(err)
		return fmt.Errorf("ERROR: %v", err)
	}
	return nil
}
