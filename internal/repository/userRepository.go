package repository

import (
	"blog/internal/models"
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type UserRepository struct {
	db           *gorm.DB
	redisSession *redis.Client
	redisCode    *redis.Client
	ctx          context.Context
}

func NewUserRepository(db *gorm.DB, redisSession, redisCode *redis.Client) *UserRepository {
	return &UserRepository{
		db:           db,
		redisSession: redisSession,
		redisCode:    redisCode,
		ctx:          context.Background(),
	}
}

func (u *UserRepository) CreateUser(user *models.User) error {
	return u.db.Create(user).Error
}

func (u *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := u.db.Where("Email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) UpdateUser(user *models.User) error {
	return u.db.Model(user).Updates(user).Error
}

func (u *UserRepository) CreateCode(email, code string) error {
	return u.redisCode.Set(u.ctx, email, code, 10*time.Minute).Err()
}

func (u *UserRepository) GetCodeByEmail(email string) (string, error) {
	code, err := u.redisCode.Get(u.ctx, email).Result()
	if err != nil {
		return "", err
	}
	return code, nil
}

func (u *UserRepository) CreateSessionID(sessionID, userID string) error {
	return u.redisSession.Set(u.ctx, sessionID, userID, 24*time.Hour).Err()
}

func (u *UserRepository) GetUserIdBySession(sessionID string) (string, error) {
	userID, err := u.redisSession.Get(u.ctx, sessionID).Result()
	if err != nil {
		return "", err
	}
	return userID, nil
}
