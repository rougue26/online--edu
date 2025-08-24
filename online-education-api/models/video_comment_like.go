package models

import (
	"time"
)

// VideoCommentLike 视频评论点赞模型
type VideoCommentLike struct {
	ID        int       `json:"id"`
	CommentID int       `json:"comment_id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}