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

// Returns true if user has liked the post, and false if not
func HasUserLikedPost(userId int, postId int) (bool, error) {

	var count int

	row := db.QueryRow(""+
		"SELECT COUNT(l.UserId) "+
		"FROM Post p "+
		"JOIN `Like` l ON l.PostId = p.PostId "+
		"WHERE p.PostId = ? AND l.UserId = ?;", postId, userId)
	err := row.Scan(&count)
	if errors.Is(err, sql.ErrNoRows) {
		return false, fmt.Errorf("returned no rows when there should always be one")
	}
	if err != nil {
		log.Println(err)
		return false, err
	}

	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

// Gets the total number of likes a post has received
func GetLikeCountForPost(postId int) (int, error) {

	var totalLikes int

	row := db.QueryRow(""+
		"SELECT COUNT(l.UserId) "+
		"FROM Post p "+
		"JOIN `Like` l ON l.PostId = p.PostId "+
		"WHERE p.PostId = ?", postId)
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
