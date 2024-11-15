package database

import (
	"database/sql"
	"errors"
	"log"
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
	row := db.QueryRow("SELECT PostId, UserId, BasisCanvasId, Image, Caption, TimeSpentDrawing, CreationDateTime, Active FROM Post WHERE PostId = ?", postId)

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
// Hides inactive posts
func GetPostsWithLikeCountSortedByMostLikes(limit int) ([]models.PostExtended, error) {
	query := `
    SELECT p.PostId, p.UserId, p.BasisCanvasId, p.Image, p.Caption, p.TimeSpentDrawing, p.CreationDateTime, COUNT(l.PostId) AS LikeCount
    FROM Post p
    LEFT JOIN rapidart.Like l ON p.PostId = l.PostId
	WHERE p.Active = true
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

// Fetches recent posts. Use the activeOnes parameter to specify whether you want active or inactive posts.
//
// If no posts are found, an empty slice is returned.
//
// Returns: list of PostExtended with likecount, error
func GetRecentPostsWithLikes(limit int, activeOnes bool) ([]models.PostExtended, error) {
	query := "SELECT p.PostId, p.UserId, p.BasisCanvasId, p.Image, p.Caption, p.TimeSpentDrawing, p.CreationDateTime, p.Active, COUNT(l.PostId) AS LikeCount" +
		" FROM Post p" +
		" LEFT JOIN `Like` l ON p.PostId = l.PostId" +
		" WHERE p.Active = ?" +
		" GROUP BY p.PostId" +
		" ORDER BY p.CreationDateTime DESC" +
		" LIMIT ?"

	// Execute the query
	rows, err := db.Query(query, activeOnes, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Slice to store the results
	var posts []models.PostExtended

	// Iterate through the rows
	for rows.Next() {
		var post models.PostExtended
		err := rows.Scan(&post.PostId, &post.UserId, &post.BasisCanvasId, &post.Image, &post.Caption, &post.TimeSpentDrawing, &post.CreationDateTime, &post.Active, &post.LikeCount)
		if err != nil {
			return nil, err
		}

		// Convert CreationDateTime to local time
		post.CreationDateTime = post.CreationDateTime.Local()

		posts = append(posts, post)
	}
	return posts, nil
}

// Fetches recent posts by a users follows.
// Use the activeOnes parameter to specify whether you want active or inactive posts.
// The userId is the user whos follows' posts are going to be returned.
//
// This query is a combination of GetRecentPostsWithLikes and GetPostsByUserId, and was made to improve performance.
//
// If no posts are found, an empty slice is returned.
//
// Returns: list of PostExtended with likecount, error
func GetUsersFollowsRecentPostsWithLikes(userId int, limit int, activeOnes bool) ([]models.PostExtended, error) {
	query := "SELECT p.PostId, p.UserId, p.BasisCanvasId, p.Image, p.Caption, p.TimeSpentDrawing, p.CreationDateTime, p.Active, COUNT(l.PostId) AS LikeCount" +
		" FROM `User` u" +
		" JOIN Follow f ON u.UserId = f.FollowerUserId" +
		" JOIN Post p ON f.FolloweeUserId = p.UserId" +
		" LEFT JOIN `Like` l ON p.PostId = l.PostId" +
		" WHERE u.UserId = ? AND p.Active = ?" +
		" GROUP BY p.PostId" +
		" ORDER BY p.CreationDateTime DESC" +
		" LIMIT ?"

	// Execute the query
	rows, err := db.Query(query, userId, activeOnes, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Slice to store the results
	var posts []models.PostExtended

	// Iterate through the rows
	for rows.Next() {
		var post models.PostExtended
		err := rows.Scan(&post.PostId, &post.UserId, &post.BasisCanvasId, &post.Image, &post.Caption, &post.TimeSpentDrawing, &post.CreationDateTime, &post.Active, &post.LikeCount)
		if err != nil {
			return nil, err
		}

		// Convert CreationDateTime to local time
		post.CreationDateTime = post.CreationDateTime.Local()

		posts = append(posts, post)
	}
	return posts, nil
}

// Fetches a list of posts that was created by the given user id
//
// OrderBy can take the following values: "likeCountDesc", "likeCountAsc", "creationDateTimeAsc", "creationDateTimeDesc"
//
// The following fields are populated in PostExtended: Post, LikeCount.
// If no posts are found, an empty slice is returned.
//
// NOTE: Inactive posts are hidden
func GetPostsByUserId(userId int, orderBy string, limit uint) ([]models.PostExtended, error) {

	ordering := ""

	switch orderBy {
	case "likeCountDesc":
		ordering = "LikeCount DESC"
	case "likeCountAsc":
		ordering = "LikeCount ASC"
	case "creationDateTimeAsc":
		ordering = "p.CreationDateTime Asc"
	case "creationDateTimeDesc":
		ordering = "p.CreationDateTime DESC"
	default:
		return nil, errors.New("invalid orderBy value")
	}

	query := "" +
		"SELECT p.PostId, p.UserId, p.BasisCanvasId, p.Image, p.Caption, p.TimeSpentDrawing, p.CreationDateTime, p.Active, COUNT(l.PostId) AS LikeCount " +
		"FROM `Post` p " +
		"LEFT OUTER JOIN `Like` l ON p.PostId = l.PostId " +
		"WHERE p.UserId = ? AND p.Active = true " +
		"GROUP BY p.PostId " +
		"ORDER BY " + ordering + " " +
		"LIMIT ?;"

	// Execute the query
	rows, err := db.Query(query, userId, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Slice to store the results
	var posts []models.PostExtended

	// Iterate through the rows
	for rows.Next() {
		post := models.PostExtended{}
		err := rows.Scan(&post.PostId, &post.UserId, &post.BasisCanvasId, &post.Image, &post.Caption, &post.TimeSpentDrawing, &post.CreationDateTime, &post.Active, &post.LikeCount)
		if err != nil {
			return nil, err
		}

		// Convert CreationDateTime to local time
		post.CreationDateTime = post.CreationDateTime.Local()

		posts = append(posts, post)
	}

	return posts, nil
}

func DeactivateActivePost(postId int) error {
	sqlStatement := "UPDATE Post SET Active = 0 WHERE PostId = ?"

	_, err := db.Exec(sqlStatement, postId)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
