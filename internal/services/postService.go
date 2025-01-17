package services

import (
	"blog/internal/models"
	"blog/internal/repository"
	"errors"
	"log"
	"strconv"

	"github.com/google/uuid"
)

type PostService struct {
	PostRepository *repository.PostRepository
}

func NewPostService(postRepository *repository.PostRepository) *PostService {
	return &PostService{PostRepository: postRepository}
}

// This method creates a new post.
// It sets the user ID in the post and saves it to the repository.
// It returns an error if the post creation fails.
func (p *PostService) NewPost(post *models.Post, userID uuid.UUID) error {

	post.UserID = userID

	err := p.PostRepository.CreatePost(post)
	if err != nil {
		log.Printf("Failed to create post for user %s: %v", userID.String(), err)
		return errors.New("failed to create post " + err.Error())
	}

	log.Printf("Successfully created post for user %s", userID.String())
	return nil
}

// This method retrieves all posts for the specified user.
// It converts the user's string ID to UUID and fetches the posts associated with that user.
// It returns an error if the user ID conversion fails or if fetching posts fails.
func (p *PostService) GetPosts(userIDstr string) ([]models.Post, error) {

	userID, err := uuid.Parse(userIDstr)
	if err != nil {
		log.Printf("Invalid user ID %s: %v", userIDstr, err)
		return nil, errors.New("invalid user ID " + err.Error())
	}

	posts, err := p.PostRepository.GetPosts(userID)
	if err != nil {
		log.Printf("Failed to retrieve posts for user %s: %v", userID.String(), err)
		return nil, errors.New("failed to get posts " + err.Error())
	}

	log.Printf("Successfully retrieved posts for user %s", userID.String())
	return posts, nil
}

// This method deletes a post with the specified ID.
// It converts the post's string ID to a number and calls the repository to delete the post.
// It returns an error if the post ID is invalid or if the post deletion fails.
func (p *PostService) DeletePost(postIDStr string, userID uuid.UUID) error {

	postID, err := strconv.ParseUint(postIDStr, 10, 64)
	if err != nil {
		log.Printf("Invalid post ID %s: %v", postIDStr, err)
		return errors.New("invalid post Id" + err.Error())
	}

	err = p.PostRepository.DeletePost(uint(postID), userID)
	if err != nil {
		log.Printf("Failed to delete post %s for user %s: %v", postIDStr, userID.String(), err)
		return errors.New("failed to delete post" + err.Error())
	}

	log.Printf("Successfully deleted post %s for user %s", postIDStr, userID.String())
	return nil
}
