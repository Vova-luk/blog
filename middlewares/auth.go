package middlewares

import (
	"blog/internal/repository"
	"context"
	"net/http"

	"github.com/google/uuid"
)

const userIDKey string = "userID"

// SessionMiddleware is middleware for processing user sessions.
// It retrieves the session ID from the "sessionID" cookie, gets the userID from the repository,
// checks it for validity and adds userID to the request context.
// If the session is invalid or there is a data error, returns a 401 Unauthorized error.
// If the check is successful, passes the request to the next handler with the updated context.
func SessionMiddleware(userRepository *repository.UserRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			session, err := r.Cookie("sessionID")
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			userIDStr, err := userRepository.GetUserIdBySession(session.Value)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			userID, err := uuid.Parse(userIDStr)
			if err != nil {
				http.Error(w, "Invalid session data", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), userIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
