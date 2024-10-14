package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"rapidart/internal/models"
	"time"
)

/*
// Saves a post to the database

	func SavePost(userId, basisCanvasId int, image []byte, caption string, timeSpent int) error {
		// query to insert new post "?" are placeholder
		query := `INSERT INTO Post (UserId, BasisCanvasId, Image, Caption, TimeSpentDrawing, CreationDateTime, Active)
	              VALUES (?, ?, ?, ?, ?, ?, ?)`

		// execute the query with creationdatetime as timenow and active as true
		_, err := db.Exec(query, userId, basisCanvasId, image, caption, timeSpent, time.Now(), true)
		if err != nil {
			return err
		}

		return nil
	}
*/
func NewPost(newPostModel models.Post) error {
	_, err := UserById(newPostModel.UserID)
	if err != nil {
		log.Println("user does not exist")
		fmt.Println(err)
		return fmt.Errorf("ERROR: %v", err)
	}

	newPostModel.CreationTimestamp = time.Now()

	sqlInsert := `
		INSERT INTO Post (
		                  PostId,
		                  UserId,
		                  BasisCanvasId,
		                  Image,
		                  Caption,
		                  TimeSpentDrawing,
		                  CreationTimeStamp,
		                  Active
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?);`

	_, err = db.Exec(sqlInsert,
		newPostModel.PostID,
		newPostModel.UserID,
		newPostModel.BasisCanvasID,
		newPostModel.Image,
		newPostModel.Caption,
		newPostModel.TimeSpentDrawing,
		newPostModel.CreationTimestamp,
		newPostModel.Active,
	)
	if err != nil {
		log.Println("Error: ", err)
		fmt.Println(err)
		return fmt.Errorf("ERROR: %v", err)
	}

	return nil
}

// Gets a post by specified ID
func GetPostById(postId int) (models.Post, error) {
	var post models.Post

	// query to select post with specified PostID
	row := db.QueryRow("SELECT PostId, UserId, BasisCanvasId, Image, Caption, TimeSpentDrawing, CreationDateTime FROM Post WHERE PostId = ?", postId)

	// scan the row into fields of Post struct
	err := row.Scan(&post.PostID, &post.UserID, &post.BasisCanvasID, &post.Image, &post.Caption, &post.TimeSpentDrawing, &post.CreationTimestamp)

	if errors.Is(err, sql.ErrNoRows) {
		return post, errors.New("no post found with that id")
	} // no row
	if err != nil {
		return post, err
	} // other errors

	post.CreationTimestamp = post.CreationTimestamp.Local()

	// returns post struct with data and no error
	return post, nil
}
