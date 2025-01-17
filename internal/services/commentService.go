package services

import (
	"blog/internal/models"
	"blog/internal/repository"
	"errors"
	"log"
	"strconv"

	"github.com/google/uuid"
)

type CommentServices struct {
	CommentRepository *repository.CommentRepository
}

func NewCommentService(commentRepository *repository.CommentRepository) *CommentServices {
	return &CommentServices{CommentRepository: commentRepository}
}

// This method creates a new comment for the specified post.
// It sets the user and post IDs, and then saves the comment to the repository.
// It returns an error if the post ID is invalid or if the comment creation fails.
func (c *CommentServices) CreateComment(comment *models.Comment, userID uuid.UUID, postIDstr string) error {

	comment.UserID = userID

	postID, err := strconv.ParseUint(postIDstr, 10, 32)
	if err != nil {
		log.Printf("Invalid post ID %s: %v", postIDstr, err)
		return errors.New("invalid post ID" + err.Error())
	}
	comment.PostId = uint(postID)

	if err := c.CommentRepository.CreateComment(comment); err != nil {
		log.Printf("Failed to create comment for post %s: %v", postIDstr, err)
		return errors.New("failed to create comment" + err.Error())
	}

	log.Printf("Successfully created comment for post %s by user %s", postIDstr, userID.String())
	return nil
}

// This method retrieves all comments for the specified post.
// It converts the post's string ID to a number and fetches the comments associated with it.
// It returns an error if the post ID is invalid or if fetching comments fails.
func (c *CommentServices) GetComments(postIDstr string) ([]models.Comment, error) {

	postID, err := strconv.ParseUint(postIDstr, 10, 32)
	if err != nil {
		log.Printf("Invalid post ID %s: %v", postIDstr, err)
		return nil, errors.New("invalid post ID" + err.Error())
	}

	comments, err := c.CommentRepository.GetCommentsByPostId(uint(postID))
	if err != nil {
		log.Printf("Failed to get comments for post %s: %v", postIDstr, err)
		return nil, errors.New("failed to get comments" + err.Error())
	}

	log.Printf("Successfully retrieved comments for post %s", postIDstr)
	return comments, nil
}

// This method deletes a comment with the specified ID for the given post.
// It converts the comment and post IDs from strings to numbers and calls the repository to delete the comment.
// It returns an error if the IDs are invalid or if the comment deletion fails.
func (c *CommentServices) DeleteComment(commentIDstr string, postIDstr string, userID uuid.UUID) error {

	commentID, err := strconv.ParseUint(commentIDstr, 10, 32)
	if err != nil {
		log.Printf("Invalid comment ID %s: %v", commentIDstr, err)
		return errors.New("invalid comment ID" + err.Error())
	}

	postID, err := strconv.ParseUint(postIDstr, 10, 32)
	if err != nil {
		log.Printf("Invalid post ID %s: %v", postIDstr, err)
		return errors.New("invalid post ID" + err.Error())
	}

	if err := c.CommentRepository.DeleteComment(uint(commentID), uint(postID), userID); err != nil {
		log.Printf("Failed to delete comment %s for post %s by user %s: %v", commentIDstr, postIDstr, userID.String(), err)
		return errors.New("failed to delete comment" + err.Error())
	}

	log.Printf("Successfully deleted comment %s for post %s by user %s", commentIDstr, postIDstr, userID.String())
	return nil
}
