package comment

import (
	"errors"
	"rapidart/internal/database"
	"rapidart/internal/models"
	"time"
)

// Validates and adds comment to post.
func CommentPost(postId int, userId int, message string) (int, error) {

	// Validate
	if len(message) > 512 || len(message) < 1 {
		return 0, errors.New("message-invalid-format")
	}

	// Create comment
	comment := models.Comment{
		UserId:           userId,
		PostId:           postId,
		Message:          message,
		CreationDateTime: time.Now(),
	}

	return database.AddCommentToPost(comment)
}

// Get comments by post id
func GetCommentsByPostId(postId int) ([]models.Comment, error) {
	return database.GetAllCommentsFromPost(postId)
}

// Get comments with commenter by a post id
func GetCommentsWithCommenterByPostId(postId int) ([]models.CommentExtended, error) {
	var commentsExt []models.CommentExtended

	// Get comments by post id
	comments, err := database.GetAllCommentsFromPost(postId)
	if err != nil {
		return []models.CommentExtended{}, err
	}

	// Loop through and fill in commenters
	for _, c := range comments {
		user, err := database.GetUserById(c.UserId)
		if err != nil {
			return []models.CommentExtended{}, err
		}

		commentsExt = append(commentsExt, models.CommentExtended{
			Comment:   c,
			Commenter: user,
		})
	}

	return commentsExt, nil
}
