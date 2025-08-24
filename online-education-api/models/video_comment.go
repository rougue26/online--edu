package models

import (
	"time"
)

// VideoComment 视频评论模型
type VideoComment struct {
	ID        int       `json:"id"`
	VideoID   int       `json:"video_id"`
	UserID    int64     `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// VideoCommentWithUserInfo 带用户信息的评论结构
type VideoCommentWithUserInfo struct {
	VideoComment
	AuthorName string `json:"authorName"`
	VideoTitle string `json:"videoTitle"`
}