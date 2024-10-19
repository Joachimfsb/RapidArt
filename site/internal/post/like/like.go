package post

import (
	"log"
	"rapidart/internal/database"
	"rapidart/internal/models"
)

func GetTopLikedPosts() ([]models.Post, error) {
	// Fetch raw rows from the database layer
	rows, err := database.GetPostsWithLikeCounts()
	if err != nil {
		log.Println("Error fetching posts with like counts:", err)
		return nil, err
	}
	defer rows.Close()

	// Process rows into Post structs
	var topPosts []models.Post
	for rows.Next() {
		var post models.Post
		// Scan the row into the Post model, including the like count
		err := rows.Scan(&post.PostId, &post.UserId, &post.BasisCanvasId, &post.Image, &post.Caption, &post.TimeSpentDrawing, &post.CreationDateTime, &post.LikeCount)
		if err != nil {
			log.Println("Error scanning post data:", err)
			continue
		}
		topPosts = append(topPosts, post)
	}

	// Return the processed list of top posts
	return topPosts, nil
}
