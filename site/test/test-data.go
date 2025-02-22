// This package contains test related helper functions
package test

import (
	"rapidart/internal/crypto"
	"rapidart/internal/models"
	"strconv"
	"time"
)

var userCount = 0

// Generate a test user
func GenTestUser() models.User {

	userCount++

	salt := crypto.GenerateRandomCharacters(5)

	return models.User{
		UserId:       userCount,                                      // Unique
		Username:     "test" + strconv.Itoa(userCount),               // Unique
		Email:        "test" + strconv.Itoa(userCount) + "@test.com", // Unique
		Displayname:  "Test testesen",
		Password:     crypto.PBDKF2("test", salt),
		PasswordSalt: salt,
		CreationTime: time.Now(),
		Role:         "user",
		Bio:          "My name is user!",
		Profilepic:   nil,
	}
}

var postCount = 0

// Generate a test post
func GenTestPost(userId int, basisCanvasId int, extended bool) (models.Post, models.PostExtended) {

	postCount++

	if extended {
		return models.Post{}, models.PostExtended{
			Post: models.Post{
				PostId:           postCount,
				UserId:           userId,
				BasisCanvasId:    basisCanvasId,
				Image:            nil,
				Caption:          "Test caption",
				TimeSpentDrawing: 180,
				CreationDateTime: time.Now(),
				Active:           true,
			},
			LikeCount: 1,
		}
	} else {
		return models.Post{
			PostId:           postCount,
			UserId:           userId,
			BasisCanvasId:    basisCanvasId,
			Image:            nil,
			Caption:          "Test caption",
			TimeSpentDrawing: 180,
			CreationDateTime: time.Now(),
			Active:           true,
		}, models.PostExtended{}
	}
}

var basisGalleryCount = 0

// Generate a basis gallery
func GenBasisGallery() models.BasisGallery {

	basisGalleryCount++

	return models.BasisGallery{
		BasisGalleryId: basisGalleryCount,
		StartDateTime:  time.Now(),
		EndDateTime:    time.Now().AddDate(0, 0, 1),
	}
}

var basisCanvasCount = 0

func GenBasisCanvas(basisGalleryId int) models.BasisCanvas {

	basisCanvasCount++

	return models.BasisCanvas{
		BasisCanvasId:  basisCanvasCount,
		BasisGalleryId: basisGalleryId,
		Type:           "test type",
		Image:          nil,
	}
}

func GenLike(userId int, postId int) models.Like {

	return models.Like{
		UserId: userId,
		PostId: postId,
	}
}

var commentCount = 0

func GenComment(userId int, postId int) models.Comment {

	commentCount++

	return models.Comment{
		CommentId:        commentCount,
		UserId:           userId,
		PostId:           postId,
		Message:          "Wow, amazing stuff",
		CreationDateTime: time.Now(),
	}
}

func GenReport(userId int, postId int) models.Report {

	return models.Report{
		UserId:           userId,
		PostId:           postId,
		Message:          "Innapropriate drawing!!",
		CreationDateTime: time.Now(),
	}
}

func GenFollow(followerId int, followeeId int) models.Follow {

	return models.Follow{
		FollowerUserId: followerId,
		FolloweeUserId: followeeId,
	}
}
