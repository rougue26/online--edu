package models

import (
	"time"
)

// Payment 支付记录模型
type Payment struct {
	ID            int64     `json:"id"`
	OrderID       string    `json:"order_id"`
	UserID        int64     `json:"user_id"`
	CourseID      int64     `json:"course_id"`
	Amount        float64   `json:"amount"`
	PaymentMethod string    `json:"payment_method"` // wechat, alipay
	Status        string    `json:"status"`         // pending, completed, failed, refunded
	TransactionID string    `json:"transaction_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// CreatePaymentRequest 创建支付请求
type CreatePaymentRequest struct {
	CourseID      int64  `json:"course_id"`
	PaymentMethod string `json:"payment_method"` // wechat, alipay
}

// PaymentResponse 支付响应
type PaymentResponse struct {
	ID            int64     `json:"id"`
	OrderID       string    `json:"order_id"`
	CourseID      int64     `json:"course_id"`
	Amount        float64   `json:"amount"`
	PaymentMethod string    `json:"payment_method"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	CourseTitle   string    `json:"course_title"`
}