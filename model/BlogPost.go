package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlogPostStatus string

const (
	Draft     BlogPostStatus = "DRAFT"
	Published BlogPostStatus = "PUBLISHED"
	Closed    BlogPostStatus = "CLOSED"
	Active    BlogPostStatus = "ACTIVE"
	Famous    BlogPostStatus = "FAMOUS"
)

type BlogPost struct {
	ID           uuid.UUID         `gorm:"primaryKey" json:"id,omitempty"`
	AuthorID     int               `gorm:"not null"`
	TourID       int               `gorm:"not null"`
	Title        string            `gorm:"not null"`
	Description  string            `gorm:"not null"`
	CreationDate time.Time         `gorm:"not null"`
	Images       []string          `gorm:"type:jsonb;default:'[]';serializer:json" json:"images"`
	Comments     []BlogPostComment `gorm:"type:jsonb;default:'[]';serializer:json" json:"comments"`
	Ratings      []BlogPostRating  `gorm:"type:jsonb;default:'[]';serializer:json" json:"ratings"`
	Status       BlogPostStatus    `gorm:"not null"`
}

func (blogPost *BlogPost) BeforeCreate(scope *gorm.DB) error {
	blogPost.ID = uuid.New()
	return nil
}

type BlogPostComment struct {
	Text            string    `json:"text"`
	CreationTime    time.Time `json:"creation_time"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

type BlogPostRating struct {
	CreationTime time.Time `json:"creation_time"`
	IsPositive   bool      `json:"is_positive"`
}
