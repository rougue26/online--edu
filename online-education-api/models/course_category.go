package models

import (
	"time"
)

// CourseCategory 课程分类模型
type CourseCategory struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	ParentID  *int64    `json:"parent_id,omitempty"` // 父分类ID，可为空
	SortOrder int       `json:"sort_order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Children  []*CourseCategory `json:"children,omitempty"` // 子分类
}

// CourseCategoryResponse 课程分类响应模型
type CourseCategoryResponse struct {
	ID        int64                  `json:"id"`
	Name      string                 `json:"name"`
	ParentID  *int64                 `json:"parent_id,omitempty"`
	SortOrder int                    `json:"sort_order"`
	CreatedAt time.Time              `json:"created_at"`
	Children  []*CourseCategoryResponse `json:"children,omitempty"`
}