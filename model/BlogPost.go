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
	ID             uuid.UUID         `gorm:"primaryKey" json:"id,omitempty"`
	AuthorID       string            `gorm:"not null"`
	AuthorUsername string            `gorm:"-" json:"authorUsername"`
	TourID         int               `gorm:"not null"`
	Title          string            `gorm:"not null" json:"title"`
	Description    string            `gorm:"not null" json:"description"`
	CreationDate   time.Time         `gorm:"not null" json:"creationDate"`
	Images         []string          `gorm:"type:jsonb;default:'[]';serializer:json" json:"imageURLs"`
	Comments       []BlogPostComment `gorm:"type:jsonb;default:'[]';serializer:json" json:"comments"`
	Ratings        []BlogPostRating  `gorm:"type:jsonb;default:'[]';serializer:json" json:"ratings"`
	Status         BlogPostStatus    `gorm:"not null" json:"status"`
}

type BlogPostComment struct {
	AuthorID        string    `json:number`
	AuthorUsername  string    `json:"authorUsername"`
	Text            string    `json:"text"`
	CreationTime    time.Time `json:"creationTime"`
	LastUpdatedTime time.Time `json:"lastUpdatedTime"`
}

type BlogPostRating struct {
	AuthorID       string    `json:number`
	AuthorUsername string    `json:text`
	CreationTime   time.Time `json:"creation_time"`
	IsPositive     bool      `json:"is_positive"`
}

type User struct {
	ID       string `json:id"`
	Username string `json:"username"`
}

type PagedResult[T any] struct {
	Results    []T `json:"results"`
	TotalCount int `json:"totalCount"`
}

func (blogPost *BlogPost) BeforeCreate(scope *gorm.DB) error {
	blogPost.ID = uuid.New()
	return nil
}
