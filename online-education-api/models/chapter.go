package models

import (
	"time"
)

// Chapter 章节模型
type Chapter struct {
	ID        int64     `json:"id"`
	CourseID  int64     `json:"course_id"`
	Title     string    `json:"title"`
	SortOrder int       `json:"sort_order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Lessons   []*Lesson `json:"lessons,omitempty"`
}

// ChapterResponse 章节响应模型
type ChapterResponse struct {
	ID        int64          `json:"id"`
	Title     string         `json:"title"`
	SortOrder int            `json:"sort_order"`
	Lessons   []*Lesson      `json:"lessons,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
}

// CreateChapterRequest 创建章节请求
type CreateChapterRequest struct {
	CourseID  int64  `json:"course_id" binding:"required"`
	Title     string `json:"title" binding:"required,min=2,max=100"`
	SortOrder int    `json:"sort_order"`
}

// UpdateChapterRequest 更新章节请求
type UpdateChapterRequest struct {
	Title     string `json:"title" min=2,max=100`
	SortOrder int    `json:"sort_order"`
}