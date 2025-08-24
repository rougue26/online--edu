package models

import (
	"time"
)

// UserCourse 用户课程关系模型
type UserCourse struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	CourseID  int64     `json:"course_id"`
	OrderID   string    `json:"order_id"`
	Price     float64   `json:"price"`
	Status    int       `json:"status"` // 1: 已购买, 0: 已退款
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserCourseResponse 用户课程响应模型
type UserCourseResponse struct {
	ID        int64      `json:"id"`
	CourseID  int64      `json:"course_id"`
	Course    *Course    `json:"course,omitempty"`
	Price     float64    `json:"price"`
	Status    int        `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
}

// EnrollCourseRequest 报名课程请求
type EnrollCourseRequest struct {
	CourseID int64 `json:"course_id" binding:"required"`
}