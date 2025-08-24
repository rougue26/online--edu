package models

import (
	"time"
)

// Course 课程模型
type Course struct {
	ID            int64     `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	CoverImage    string    `json:"cover_image"`
	Price         float64   `json:"price"`
	OriginalPrice float64   `json:"original_price"`
	CategoryID    int64     `json:"category_id"`
	TeacherID     int64     `json:"teacher_id"`
	Level         int       `json:"level"` // 1: 初级, 2: 中级, 3: 高级
	Duration      int       `json:"duration"` // 课程时长(分钟)
	StudentCount  int       `json:"student_count"`
	Rating        float64   `json:"rating"`
	Status        int       `json:"status"` // 1: 上架, 0: 下架
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// CourseResponse 课程响应模型
type CourseResponse struct {
	ID            int64      `json:"id"`
	Title         string     `json:"title"`
	Description   string     `json:"description"`
	CoverImage    string     `json:"cover_image"`
	Price         float64    `json:"price"`
	OriginalPrice float64    `json:"original_price"`
	CategoryID    int64      `json:"category_id"`
	TeacherID     int64      `json:"teacher_id"`
	CategoryName  string     `json:"category_name"`
	TeacherName   string     `json:"teacher_name"`
	Level         int        `json:"level"`
	Duration      int        `json:"duration"`
	StudentCount  int        `json:"student_count"`
	Rating        float64    `json:"rating"`
	Status        int        `json:"status"`
	CreatedAt     time.Time  `json:"created_at"`
}

// CourseDetailResponse 课程详情响应模型
type CourseDetailResponse struct {
	Course
	Category    *CourseCategoryResponse `json:"category"`
	Teacher     *UserResponse   `json:"teacher"`
	Chapters    []*Chapter      `json:"chapters"`
	TotalLessons int            `json:"total_lessons"`
}

// CreateCourseRequest 创建课程请求
type CreateCourseRequest struct {
	Title         string  `json:"title" binding:"required,min=2,max=100"`
	Description   string  `json:"description"`
	CoverImage    string  `json:"cover_image"`
	Price         float64 `json:"price"`
	OriginalPrice float64 `json:"original_price"`
	CategoryID    int64   `json:"category_id" binding:"required"`
	Level         int     `json:"level" binding:"required,oneof=1 2 3"`
	Status        int     `json:"status" binding:"oneof=0 1"`
}

// UpdateCourseRequest 更新课程请求
type UpdateCourseRequest struct {
	Title         string  `json:"title" min=2,max=100`
	Description   string  `json:"description"`
	CoverImage    string  `json:"cover_image"`
	Price         float64 `json:"price"`
	OriginalPrice float64 `json:"original_price"`
	CategoryID    int64   `json:"category_id"`
	Level         int     `json:"level" oneof=1 2 3`
	Status        int     `json:"status" oneof=0 1`
}