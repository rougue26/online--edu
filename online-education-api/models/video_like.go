package models

import (
	"time"
)

// VideoLike 视频点赞模型
type VideoLike struct {
	ID        int       `json:"id"`
	VideoID   int       `json:"video_id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}