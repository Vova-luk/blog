package main

import (
	"blog/db"
	"blog/internal/handlers"
	"blog/internal/models"
	"blog/internal/repository"
	"blog/internal/services"
	"blog/middlewares"
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"net/http"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Connecting to PostgreSQL
	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Bad connection to PostgreSQL: %v", err)
	}

	if err := database.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}); err != nil {
		log.Fatalf("Bad migration: %v", err)
	}

	// Connecting to Redis
	redisSession, err := db.ConnectToRedis(0)
	if err != nil {
		log.Fatalf("Bad connection to Redis: %v", err)
	}

	// Connecting to Redis
	redisCode, err := db.ConnectToRedis(1)
	if err != nil {
		log.Fatalf("Bad connection to Redis: %v", err)
	}

	s := chi.NewRouter()

	//Router for working with the user (registration, email confirmation, login)
	userRepo := repository.NewUserRepository(database, redisSession, redisCode)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	s.Post("/users", userHandler.RegisterUser)
	s.Post("/verify", userHandler.VerifyEmail)
	s.Post("/login", userHandler.LoginUser)

	//Router for working with posts (creating, receiving and deleting)
	postRepo := repository.NewPostRepository(database)
	postService := services.NewPostService(postRepo)
	postHandler := handlers.NewPostHandlers(postService)

	//Grouping routes for posts using middleware to check sessions.
	s.Group(func(s chi.Router) {
		s.Use(middlewares.SessionMiddleware(userRepo))
		s.Post("/posts", postHandler.NewPost)
		s.Delete("/posts/{postID}", postHandler.DeletePost)
	})
	s.Get("/posts/{userID}", postHandler.GetPosts)

	//Router for working with comments (creating, receiving and deleting)
	commentRepo := repository.NewCommentRepository(database)
	commentService := services.NewCommentService(commentRepo)
	commentHandler := handlers.NewCommentHandler(commentService)

	//Grouping routes for comments using middleware to check sessions.
	s.Group(func(s chi.Router) {
		s.Use(middlewares.SessionMiddleware(userRepo))
		s.Post("/posts/{postID}/comment", commentHandler.NewComment)
		s.Delete("/posts/{postID}/comment/{commentID}", commentHandler.DeleteComment)
	})
	s.Get("/posts/{postID}/comment", commentHandler.GetComments)

	http.ListenAndServe(":8080", s)

}
