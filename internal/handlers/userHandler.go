package handlers

import (
	"blog/internal/models"
	"blog/internal/services"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

// This handler handles user registration by checking the data from the request,
// and if registration is successful, it returns status 201 (Created).
// If the data is incorrect or an error occurs during registration, appropriate errors are returned.
func (u *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = u.UserService.RegisterUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// This handler processes email verification by checking the code from the request.
// On successful verification, it returns status 204 (No Content).
// If an error occurs, an appropriate error is returned.
func (u *UserHandler) VerifyEmail(w http.ResponseWriter, r *http.Request) {

	type VerifyEmailRequest struct {
		Email string
		Code  string
	}

	var req VerifyEmailRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Invalid JSON received: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := u.UserService.VerifyEmail(req.Email, req.Code); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// This handler handles user login by checking the email and password.
// On successful authentication, status 200 (OK) is returned along with user information.
// In case of errors, appropriate error codes are returned.
func (u *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {

	type LoginRequest struct {
		Email    string
		Password string
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Invalid JSON received: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}
	defer r.Body.Close()

	user, sessionID, err := u.UserService.LoginUser(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	cookie := &http.Cookie{
		Name:     "sessionID",
		Value:    sessionID,
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(24 * time.Hour),
	}
	http.SetCookie(w, cookie)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("Failed to encode user: %v", err)
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
	}
}
