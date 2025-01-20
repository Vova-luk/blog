package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Username   string    `json:"username"`
	Email      string    `gorm:"type:varchar(255);not null;unique" json:"email"`
	Password   string    `gorm:"type:varchar(255);not null" json:"password"`
	IsVerified bool      `gorm:"default:false" json:"is_verified"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Post struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"post_id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Comment struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"comment_id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	PostId    uint      `gorm:"not null" json:"post_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
