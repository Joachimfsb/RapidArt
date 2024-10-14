package database

import (
	"database/sql"
	"errors"
	"rapidart/internal/models"
	"time"
)

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

// Gets a post by specified ID
func GetPostById(postId int) (models.Post, error) {
	var post models.Post

	// query to select post with specified PostID
	row := db.QueryRow("SELECT PostId, UserId, BasisCanvasId, Image, Caption, TimeSpentDrawing, CreationDateTime FROM Post WHERE PostId = ?", postId)

	// scan the row into fields of Post struct
	err := row.Scan(&post.PostId, &post.UserId, &post.BasisCanvasId, &post.Image, &post.Caption, &post.TimeSpentDrawing, &post.CreationDateTime)

	if errors.Is(err, sql.ErrNoRows) {
		return post, errors.New("no post found with that id")
	} // no row
	if err != nil {
		return post, err
	} // other errors

	post.CreationDateTime = post.CreationDateTime.Local()

	// returns post struct with data and no error
	return post, nil
}
