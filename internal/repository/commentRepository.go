package repository

import (
	"blog/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (c *CommentRepository) CreateComment(comment *models.Comment) error {
	return c.db.Create(comment).Error
}

func (c *CommentRepository) GetCommentsByPostId(postID uint) ([]models.Comment, error) {
	var comments []models.Comment
	err := c.db.Where("post_id = ?", postID).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, err
}

func (c *CommentRepository) DeleteComment(commentID uint, postID uint, userID uuid.UUID) error {
	return c.db.Where("id = ? AND user_id = ? AND post_id = ?", commentID, userID, postID).Delete(&models.Comment{}).Error
}
