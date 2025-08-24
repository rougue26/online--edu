package services

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"online-education-api/models"
	"time"
)

// PaymentService 支付服务接口
type PaymentService interface {
	CreatePayment(userID, courseID int64, paymentMethod string) (*models.Payment, error)
	GetPaymentByID(id int64) (*models.Payment, error)
	GetPaymentsByUserID(userID int64, page, pageSize int) ([]*models.PaymentResponse, int, error)
	UpdatePaymentStatus(orderID, transactionID, status string) error
	GetPaymentByOrderID(orderID string) (*models.Payment, error)
}

// paymentService 支付服务实现
type paymentService struct {
	db *sql.DB
}

// NewPaymentService 创建支付服务实例
func NewPaymentService(db *sql.DB) PaymentService {
	return &paymentService{db: db}
}

// CreatePayment 创建支付订单
func (s *paymentService) CreatePayment(userID, courseID int64, paymentMethod string) (*models.Payment, error) {
	// 检查课程是否存在
	courseService := NewCourseService(s.db)
	course, err := courseService.GetCourseDetail(courseID)
	if err != nil {
		return nil, errors.New("课程不存在")
	}

	// 生成订单ID (时间戳+随机数)
	now := time.Now()
	orderID := fmt.Sprintf("ORD-%d-%06d", now.Unix(), rand.Intn(1000000))

	// 创建支付记录
	payment := &models.Payment{
		OrderID:       orderID,
		UserID:        userID,
		CourseID:      courseID,
		Amount:        course.Course.Price,
		PaymentMethod: paymentMethod,
		Status:        "pending",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	// 插入数据库
	query := `INSERT INTO payments (order_id, user_id, course_id, amount, payment_method, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := s.db.Exec(query, payment.OrderID, payment.UserID, payment.CourseID, payment.Amount, payment.PaymentMethod, payment.Status, payment.CreatedAt, payment.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("创建支付订单失败: %v", err)
	}

	// 获取插入的ID
	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("获取支付订单ID失败: %v", err)
	}
	payment.ID = id

	return payment, nil
}

// GetPaymentByID 根据ID获取支付记录
func (s *paymentService) GetPaymentByID(id int64) (*models.Payment, error) {
	query := `SELECT id, order_id, user_id, course_id, amount, payment_method, status, transaction_id, created_at, updated_at FROM payments WHERE id = ?`
	row := s.db.QueryRow(query, id)

	var payment models.Payment
	if err := row.Scan(&payment.ID, &payment.OrderID, &payment.UserID, &payment.CourseID, &payment.Amount, &payment.PaymentMethod, &payment.Status, &payment.TransactionID, &payment.CreatedAt, &payment.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("支付记录不存在")
		}
		return nil, fmt.Errorf("获取支付记录失败: %v", err)
	}

	return &payment, nil
}

// GetPaymentsByUserID 根据用户ID获取支付记录列表
func (s *paymentService) GetPaymentsByUserID(userID int64, page, pageSize int) ([]*models.PaymentResponse, int, error) {
	// 计算偏移量
	offset := (page - 1) * pageSize

	// 查询支付记录
	query := `SELECT p.id, p.order_id, p.course_id, p.amount, p.payment_method, p.status, p.created_at, c.title FROM payments p LEFT JOIN courses c ON p.course_id = c.id WHERE p.user_id = ? ORDER BY p.created_at DESC LIMIT ? OFFSET ?`
	rows, err := s.db.Query(query, userID, pageSize, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("获取支付记录列表失败: %v", err)
	}
	defer rows.Close()

	// 查询总数
	var total int
	countQuery := `SELECT COUNT(*) FROM payments WHERE user_id = ?`
	if err := s.db.QueryRow(countQuery, userID).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("获取支付记录总数失败: %v", err)
	}

	// 解析结果
	var payments []*models.PaymentResponse
	for rows.Next() {
		var payment models.PaymentResponse
		if err := rows.Scan(&payment.ID, &payment.OrderID, &payment.CourseID, &payment.Amount, &payment.PaymentMethod, &payment.Status, &payment.CreatedAt, &payment.CourseTitle); err != nil {
			return nil, 0, fmt.Errorf("解析支付记录失败: %v", err)
		}
		payments = append(payments, &payment)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("读取支付记录失败: %v", err)
	}

	return payments, total, nil
}

// UpdatePaymentStatus 更新支付状态
func (s *paymentService) UpdatePaymentStatus(orderID, transactionID, status string) error {
	query := `UPDATE payments SET status = ?, transaction_id = ?, updated_at = ? WHERE order_id = ?`
	now := time.Now()

	result, err := s.db.Exec(query, status, transactionID, now, orderID)
	if err != nil {
		return fmt.Errorf("更新支付状态失败: %v", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("获取影响行数失败: %v", err)
	}

	if affected == 0 {
		return errors.New("支付订单不存在")
	}

	// 如果支付成功，创建用户课程关联
	if status == "completed" {
		payment, err := s.GetPaymentByOrderID(orderID)
		if err != nil {
			return fmt.Errorf("获取支付订单失败: %v", err)
		}

		userCourseService := NewUserCourseService(s.db)
		if err := userCourseService.EnrollCourse(payment.UserID, payment.CourseID); err != nil {
			return fmt.Errorf("创建用户课程关联失败: %v", err)
		}
	}

	return nil
}

// GetPaymentByOrderID 根据订单ID获取支付记录
func (s *paymentService) GetPaymentByOrderID(orderID string) (*models.Payment, error) {
	query := `SELECT id, order_id, user_id, course_id, amount, payment_method, status, transaction_id, created_at, updated_at FROM payments WHERE order_id = ?`
	row := s.db.QueryRow(query, orderID)

	var payment models.Payment
	if err := row.Scan(&payment.ID, &payment.OrderID, &payment.UserID, &payment.CourseID, &payment.Amount, &payment.PaymentMethod, &payment.Status, &payment.TransactionID, &payment.CreatedAt, &payment.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("支付记录不存在")
		}
		return nil, fmt.Errorf("获取支付记录失败: %v", err)
	}

	return &payment, nil
}