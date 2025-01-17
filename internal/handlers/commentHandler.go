package handlers

import (
	"blog/internal/models"
	"blog/internal/services"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type CommentHandler struct {
	CommentService *services.CommentServices
}

func NewCommentHandler(commentService *services.CommentServices) *CommentHandler {
	return &CommentHandler{CommentService: commentService}
}

// NewComment - handles the creation of a new comment for the specified post.
// It decodes the JSON request, extracts postID and userID, then calls the service method to create the comment.
// In case of errors (invalid JSON, service error), it returns the appropriate status codes.
func (c *CommentHandler) NewComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		log.Printf("Invalid JSON received: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	postIDstr := chi.URLParam(r, "postID")

	userID := r.Context().Value("userID").(uuid.UUID)
	if err := c.CommentService.CreateComment(&comment, userID, postIDstr); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetComments - handles fetching all comments for the specified post.
// It retrieves the postID from the URL parameters and calls the service method to get the comments.
// In case of an error, it returns an HTTP error if comments cannot be fetched.
func (c *CommentHandler) GetComments(w http.ResponseWriter, r *http.Request) {
	postIdstr := chi.URLParam(r, "postID")

	comments, err := c.CommentService.GetComments(postIdstr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(comments); err != nil {
		log.Printf("Failed to encode comments: %v", err)
		http.Error(w, "Failed to encode comments", http.StatusInternalServerError)
		return
	}
}

// DeleteComment - handles the deletion of a comment for the specified post.
// It retrieves the postID and commentID from the URL parameters, as well as the userID from the context.
// Then it calls the service method to delete the comment. If deletion is successful, it returns a 204 (No Content) status.
func (c *CommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	postIDstr := chi.URLParam(r, "postID")
	commentIDstr := chi.URLParam(r, "commentID")
	userID := r.Context().Value("userID").(uuid.UUID)

	if err := c.CommentService.DeleteComment(commentIDstr, postIDstr, userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
