package database

import (
	"database/sql"
	"errors"
	"rapidart/internal/models"
)

// Saves a post to the database

func AddPost(post models.Post) error {
	// query to insert new post "?" are placeholder
	query := `INSERT INTO Post (UserId, BasisCanvasId, Image, Caption, TimeSpentDrawing, CreationDateTime, Active)
	              VALUES (?, ?, ?, ?, ?, ?, ?)`

	// execute the query with creationdatetime as timenow and active as true
	_, err := db.Exec(query, post.UserId, post.BasisCanvasId, post.Image, post.Caption, post.TimeSpentDrawing, post.CreationDateTime, post.Active)
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

// Fetches posts and their like counts
func GetPostsWithLikeCountSortedByMostLikes(limit int) ([]models.PostExtended, error) {
	query := `
    SELECT p.PostId, p.UserId, p.BasisCanvasId, p.Image, p.Caption, p.TimeSpentDrawing, p.CreationDateTime, COUNT(l.PostId) AS LikeCount
    FROM Post p
    LEFT JOIN rapidart.Like l ON p.PostId = l.PostId
    GROUP BY p.PostId
    ORDER BY LikeCount DESC
    LIMIT ?;
    `

	// Execute the query
	rows, err := db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Slice to store the results
	var posts []models.PostExtended

	// Iterate through the rows
	for rows.Next() {
		var post models.PostExtended
		err := rows.Scan(&post.PostId, &post.UserId, &post.BasisCanvasId, &post.Image, &post.Caption, &post.TimeSpentDrawing, &post.CreationDateTime, &post.LikeCount)
		if err != nil {
			return nil, err
		}

		// Convert CreationDateTime to local time
		post.CreationDateTime = post.CreationDateTime.Local()

		posts = append(posts, post)
	}

	return posts, nil
}
