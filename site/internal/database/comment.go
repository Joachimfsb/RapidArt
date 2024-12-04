package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"rapidart/internal/models"
)

// Adds a new comment to a post
func AddCommentToPost(newComment models.Comment) (int, error) {
	sqlInsert := `
		INSERT INTO Comment (
		                  UserId,
		                  PostId,
		                  Message,
		                  CreationDateTime
		) VALUES (?, ?, ?, ?);`

	res, err := db.Exec(sqlInsert,
		newComment.UserId,
		newComment.PostId,
		newComment.Message,
		newComment.CreationDateTime,
	)
	if err != nil {
		log.Println("Error: ", err)
		fmt.Println(err)
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return int(id), nil
}

// Gets a list of all comments on a given post
func GetAllCommentsFromPost(postId int) ([]models.Comment, error) {
	var comments []models.Comment

	rows, err := db.Query("SELECT * FROM Comment WHERE PostId = ?", postId)
	if err != nil {
		log.Println(err)
		return []models.Comment{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment models.Comment
		err = rows.Scan(&comment.CommentId, &comment.UserId, &comment.PostId, &comment.Message, &comment.CreationDateTime)
		if err != nil {
			log.Println(err)
			return []models.Comment{}, err
		}

		comment.CreationDateTime = comment.CreationDateTime.Local()
		comments = append(comments, comment)
	}

	if errors.Is(err, sql.ErrNoRows) {
		return []models.Comment{}, fmt.Errorf("couldn't find post")
	}

	if err != nil {
		log.Println(err)
		return []models.Comment{}, err
	}

	return comments, nil
}
