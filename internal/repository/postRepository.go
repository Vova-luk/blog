package repository

import (
	"blog/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (p *PostRepository) CreatePost(post *models.Post) error {
	return p.db.Create(post).Error
}

func (p *PostRepository) GetPosts(userID uuid.UUID) ([]models.Post, error) {
	var posts []models.Post
	err := p.db.Where("user_id = ?", userID).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *PostRepository) DeletePost(postID uint, userID uuid.UUID) error {
	return p.db.Where("id = ? AND user_id = ?", postID, userID).Delete(&models.Post{}).Error
}
