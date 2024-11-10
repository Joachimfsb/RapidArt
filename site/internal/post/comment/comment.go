package comment

import (
	"rapidart/internal/database"
	"rapidart/internal/models"
)

func GetCommentsByPostId(postId int) ([]models.Comment, error) {
	return database.GetAllCommentsFromPost(postId)
}

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
