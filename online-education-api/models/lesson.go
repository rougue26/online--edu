package models

import (
	"time"
)

// Lesson 课时模型
type Lesson struct {
	ID          int64     `json:"id"`
	ChapterID   int64     `json:"chapter_id"`
	Title       string    `json:"title"`
	VideoURL    string    `json:"video_url"`
	Duration    int       `json:"duration"` // 课时时长(分钟)
	SortOrder   int       `json:"sort_order"`
	Free        int       `json:"free"` // 1: 免费, 0: 收费
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// LessonResponse 课时响应模型
type LessonResponse struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	VideoURL  string    `json:"video_url,omitempty"` // 只有购买课程或免费课时才返回视频URL
	Duration  int       `json:"duration"`
	SortOrder int       `json:"sort_order"`
	Free      int       `json:"free"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateLessonRequest 创建课时请求
type CreateLessonRequest struct {
	ChapterID  int64  `json:"chapter_id" binding:"required"`
	Title      string `json:"title" binding:"required,min=2,max=100"`
	VideoURL   string `json:"video_url" binding:"required"`
	Duration   int    `json:"duration"`
	SortOrder  int    `json:"sort_order"`
	Free       int    `json:"free" binding:"oneof=0 1"`
}

// UpdateLessonRequest 更新课时请求
type UpdateLessonRequest struct {
	Title     string `json:"title" min=2,max=100`
	VideoURL  string `json:"video_url"`
	Duration  int    `json:"duration"`
	SortOrder int    `json:"sort_order"`
	Free      int    `json:"free" oneof=0 1`
}