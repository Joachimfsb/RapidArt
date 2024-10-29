package database

import (
	"database/sql"
	"errors"
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

// Gets the total likes a user has recieved on all of their posts
func GetTotalLikesForEveryUserPost(userId int) (int, error) {

	var totalLikes int

	row := db.QueryRow(""+
		"SELECT COUNT(l.UserId) "+
		"FROM Post p "+
		"JOIN `Like` l ON l.PostId = p.PostId "+
		"WHERE p.UserId = ?", userId)
	err := row.Scan(&totalLikes)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, fmt.Errorf("returned no rows when there should always be one")
	}
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return totalLikes, nil
}
