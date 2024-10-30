package database

import (
	"database/sql"
	"errors"
	"math"
	"rapidart/internal/models"
)

// Saves a post to the database
//
// Returns: post id (-1 if fail), error
func AddPost(post models.Post) (int, error) {
	// query to insert new post "?" are placeholder
	query := `INSERT INTO Post (UserId, BasisCanvasId, Image, Caption, TimeSpentDrawing, CreationDateTime, Active)
	              VALUES (?, ?, ?, ?, ?, ?, ?)`

	// execute the query with creationdatetime as timenow and active as true
	res, err := db.Exec(query, post.UserId, post.BasisCanvasId, post.Image, post.Caption, post.TimeSpentDrawing, post.CreationDateTime, post.Active)
	if err != nil {
		return -1, err
	}

	// Get post id
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	if math.MaxInt < id {
		return -1, errors.New("int value too large") // Ran out of ints in db (requires changing types to int64 in entire program)
	}

	return int(id), nil
}

// Gets a post by specified ID
func GetPostById(postId int) (models.Post, error) {
	var post models.Post

	// query to select post with specified PostID
	row := db.QueryRow("SELECT PostId, UserId, BasisCanvasId, Image, Caption, TimeSpentDrawing, CreationDateTime FROM Post WHERE PostId = ?", postId)

	// scan the row into fields of Post struct
	err := row.Scan(&post.PostId, &post.UserId, &post.BasisCanvasId, &post.Image, &post.Caption, &post.TimeSpentDrawing, &post.CreationDateTime, &post.Active)

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

	var i = 0
	// Iterate through the rows
	for rows.Next() {
		if i < limit {

			var post models.PostExtended
			err := rows.Scan(&post.PostId, &post.UserId, &post.BasisCanvasId, &post.Image, &post.Caption, &post.TimeSpentDrawing, &post.CreationDateTime, &post.LikeCount)
			if err != nil {
				return nil, err
			}

			// Convert CreationDateTime to local time
			post.CreationDateTime = post.CreationDateTime.Local()

			posts = append(posts, post)
		} else {
			break
		}
		i += 1
	}
	return posts, nil
}

//
