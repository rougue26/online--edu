package models

import (
	"time"
)

// Post 社区帖子模型
type Post struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Category    string    `json:"category"`
	ViewCount   int       `json:"view_count"`
	CommentCount int      `json:"comment_count"`
	LikeCount   int       `json:"like_count"`
	Status      int       `json:"status"` // 1: 正常, 0: 禁用
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// PostResponse 帖子响应模型
type PostResponse struct {
	ID          int64       `json:"id"`
	UserID      int64       `json:"user_id"`
	UserName    string      `json:"user_name"`
	Avatar      string      `json:"avatar,omitempty"`
	Title       string      `json:"title"`
	Content     string      `json:"content"`
	Category    string      `json:"category"`
	ViewCount   int         `json:"view_count"`
	CommentCount int        `json:"comment_count"`
	LikeCount   int         `json:"like_count"`
	Status      int         `json:"status"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

// CreatePostRequest 创建帖子请求
type CreatePostRequest struct {
	Title   string `json:"title" binding:"required,min=2,max=100"`
	Content string `json:"content" binding:"required,min=10"`
	Category string `json:"category"`
}

// UpdatePostRequest 更新帖子请求
type UpdatePostRequest struct {
	Title   string `json:"title" min=2,max=100`
	Content string `json:"content" min=10`
	Status  int    `json:"status" oneof=0 1`
}