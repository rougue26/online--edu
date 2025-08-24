package models

import (
	"time"
)

// VideoFavorite 视频收藏模型
type VideoFavorite struct {
	ID        int       `json:"id"`
	VideoID   int       `json:"video_id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}