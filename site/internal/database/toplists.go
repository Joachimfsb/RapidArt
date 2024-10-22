package database

import (
	"log"
	"rapidart/internal/models"
)

func GetTopFollowedUsers() ([]models.User, error) {
	var topFollowed []models.User

	rows, err := db.Query("SELECT FolloweeUserId, COUNT(FollowerUserId) AS follower_count FROM Follow GROUP BY FolloweeUserId ORDER BY follower_count DESC LIMIT 10;")
	if err != nil {
		log.Println(err)
		return []models.User{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		var followedUserId int
		var count int
		err = rows.Scan(&followedUserId, &count)
		if err != nil {
			log.Println(err)
			return []models.User{}, err
		}

		user, err = GetUserById(followedUserId)
		if err != nil {
			return []models.User{}, err
		}
		topFollowed = append(topFollowed, user)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
		return []models.User{}, nil
	}

	return topFollowed, nil
}

func GetTopPosts() ([]models.Post, error) {
	var topPosts []models.Post

	rows, err := db.Query("SELECT PostId, COUNT(UserId) AS like_count FROM rapidart.Like GROUP BY PostId ORDER BY like_count DESC LIMIT 10;")
	if err != nil {
		log.Println(err)
		return []models.Post{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		var PostId int
		var count int
		err = rows.Scan(&PostId, &count)
		if err != nil {
			log.Println(err)
			return []models.Post{}, err
		}

		post, err := GetPostById(PostId)
		if err != nil {
			return []models.Post{}, err
		}
		topPosts = append(topPosts, post)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
		return []models.Post{}, nil
	}

	return topPosts, nil
}
