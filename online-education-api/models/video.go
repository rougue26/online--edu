package models

import (
	"time"
)

// Video 视频模型
type Video struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description,omitempty"`
	VideoURL       string    `json:"video_url"`
	CoverImageURL  string    `json:"cover_image_url,omitempty"`
	Duration       int       `json:"duration"`
	AuthorID       int64     `json:"author_id"`
	CategoryID     int       `json:"category_id,omitempty"`
	ViewCount      int       `json:"view_count"`
	LikeCount      int       `json:"like_count"`
	FavoriteCount  int       `json:"favorite_count"`
	IsPublic       bool      `json:"is_public"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}