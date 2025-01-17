package services

import (
	"blog/internal/models"
	"blog/internal/repository"
	"blog/utils"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

// This method handles user registration.
// It hashes the user's password, generates a verification code, and stores the user and code in the database.
// It returns an error if any of the operations fail.
func (u *UserService) RegisterUser(user *models.User) error {

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Printf("еrror while hashing password for user %s: %v", user.Email, err)
		return errors.New("error while hashing password " + err.Error())
	}

	user.Password = string(hashedPassword)

	verifyCode := utils.GenerateCode(6)
	if err := u.UserRepository.CreateCode(user.Email, verifyCode); err != nil {
		log.Printf("Error while creating verify code for user %s: %v", user.Email, err)
		return errors.New("error while creating verify code " + err.Error())
	}

	if err := u.UserRepository.CreateUser(user); err != nil {
		log.Printf("Error while creating user %s: %v", user.Email, err)
		return errors.New("error while creating user " + err.Error())
	}

	if err := utils.SendEmail(user.Email, verifyCode); err != nil {
		log.Printf("Error while sending verify code %s: %v", user.Email, err)
		return errors.New("error while sending verify code " + err.Error())
	}

	log.Printf("User %s registered successfully", user.Email)
	return nil
}

// This method handles email verification.
// It checks the verification code provided by the user and updates the user's status to "verified."
// It returns an error if the code is incorrect or if there are issues accessing the user's data.
func (u *UserService) VerifyEmail(email, code string) error {

	user, err := u.UserRepository.GetUserByEmail(email)
	if err != nil {
		log.Printf("Error while getting user by email %s: %v", email, err)
		return errors.New("error while getting user by email " + err.Error())
	}

	verifyCode, err := u.UserRepository.GetCodeByEmail(email)
	if err != nil {
		log.Printf("Error while getting verify code for email %s: %v", email, err)
		return errors.New("error while getting verify code by email " + err.Error())
	}

	if verifyCode != code {
		log.Printf("Invalid verify code for email %s", email)
		return errors.New("wrong verify code")
	}

	user.IsVerified = true

	if err := u.UserRepository.UpdateUser(user); err != nil {
		log.Printf("Error while updating user %s: %v", email, err)
		return errors.New("failed to update user " + err.Error())
	}

	log.Printf("User %s email verified successfully", email)
	return nil
}

// This method handles the user login process.
// It verifies the user's email and password, compares them with the stored data, and generates a session ID if the login is successful.
// It returns an error if the user's credentials are incorrect or if session creation fails.
func (u *UserService) LoginUser(email, password string) (*models.User, string, error) {

	user, err := u.UserRepository.GetUserByEmail(email)
	if err != nil {
		log.Printf("User %s not found: %v", email, err)
		return nil, "", errors.New("user not found")
	}

	if err := utils.CheckPasswordHash(password, user.Password); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			log.Printf("Invalid password for user %s", email)
			return nil, "", errors.New("invalid password")
		}
		log.Printf("Error comparing password and hash for user %s: %v", email, err)
		return nil, "", errors.New("error comparing password and hash")
	}

	sessionID, err := utils.GenerateSessionID()
	if err != nil {
		log.Printf("Failed to generate session ID for user %s: %v", email, err)
		return nil, "", errors.New("failed to generate session ID" + err.Error())
	}

	if err := u.UserRepository.CreateSessionID(sessionID, user.ID.String()); err != nil {
		log.Printf("Failed to create session for user %s: %v", email, err)
		return nil, "", errors.New("failed to create session ID " + err.Error())
	}

	log.Printf("User %s logged in successfully", email)
	return user, sessionID, nil
}
