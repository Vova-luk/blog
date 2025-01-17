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

type PostHandler struct {
	PostServices *services.PostService
}

func NewPostHandlers(postService *services.PostService) *PostHandler {
	return &PostHandler{
		PostServices: postService,
	}
}

// NewPost - handles the creation of a new post. It decodes the JSON request, extracts the userID from the context, and tries to create a new post through the service.
// In case of errors during decoding or creating the post, the corresponding error status is returned.
// If the post creation is successful, it returns status 201 (Created).
func (p *PostHandler) NewPost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		log.Printf("Invalid JSON received: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}
	defer r.Body.Close()

	userID := r.Context().Value("userID").(uuid.UUID)

	err := p.PostServices.NewPost(&post, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

// GetPosts - handles the request to fetch all posts for the specified user. It extracts the userID from the URL parameters and calls the service to get the posts.
// In case of an error, it returns status 500 (Internal Server Error).
// If posts are successfully retrieved, they are encoded to JSON and sent to the client with status 200 (OK).
func (p *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	userIDstr := chi.URLParam(r, "userID")

	posts, err := p.PostServices.GetPosts(userIDstr)
	if err != nil {
		http.Error(w, "Error while get posts", http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(posts); err != nil {
		log.Printf("Failed to encode posts: %v", err)
		http.Error(w, "Failed to encode posts", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

// DeletePost - handles the request to delete a post for the specified user. It extracts the postID from the URL parameters and the userID from the context.
// If the post cannot be deleted, an error with the appropriate status is returned.
// If the post is successfully deleted, status 204 (No Content).
func (p *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	postIDStr := chi.URLParam(r, "postID")
	userID := r.Context().Value("userID").(uuid.UUID)

	err := p.PostServices.DeletePost(postIDStr, userID)
	if err != nil {
		http.Error(w, "Error while delete post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
