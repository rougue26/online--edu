package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"online-education-api/models"
	"online-education-api/services"
)

// PaymentController 支付控制器
type PaymentController struct {
	paymentService services.PaymentService
}

// NewPaymentController 创建支付控制器实例
func NewPaymentController(paymentService services.PaymentService) *PaymentController {
	return &PaymentController{
		paymentService: paymentService,
	}
}

// CreatePayment 创建支付订单
func (c *PaymentController) CreatePayment(w http.ResponseWriter, r *http.Request) {
	// 从上下文获取用户ID
	userID, ok := r.Context().Value("userID").(int64)
	if !ok {
		http.Error(w, "未登录", http.StatusUnauthorized)
		return
	}

	// 解析请求参数
	var req models.CreatePaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "请求参数错误", http.StatusBadRequest)
		return
	}

	// 验证支付方式
	if req.PaymentMethod != "wechat" && req.PaymentMethod != "alipay" {
		http.Error(w, "不支持的支付方式", http.StatusBadRequest)
		return
	}

	// 创建支付订单
	payment, err := c.paymentService.CreatePayment(userID, req.CourseID, req.PaymentMethod)
	if err != nil {
		http.Error(w, "创建支付订单失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 生成支付链接或参数（这里只是示例，实际需对接微信/支付宝API）
	paymentParams := map[string]interface{}{
		"order_id": payment.OrderID,
		"amount":   payment.Amount,
		"method":   payment.PaymentMethod,
		// 实际应用中，这里会包含微信/支付宝支付所需的参数
		"pay_url": "https://api.example.com/pay/" + payment.OrderID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 200,
		"msg":  "创建支付订单成功",
		"data": paymentParams,
	})
}

// GetPaymentStatus 查询支付状态
func (c *PaymentController) GetPaymentStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["orderID"]

	// 查询支付状态
	payment, err := c.paymentService.GetPaymentByOrderID(orderID)
	if err != nil {
		http.Error(w, "查询支付状态失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 200,
		"msg":  "success",
		"data": map[string]interface{}{
			"order_id":     payment.OrderID,
			"status":       payment.Status,
			"amount":       payment.Amount,
			"payment_method": payment.PaymentMethod,
			"created_at":   payment.CreatedAt,
		},
	})
}

// GetUserPayments 获取用户支付记录
func (c *PaymentController) GetUserPayments(w http.ResponseWriter, r *http.Request) {
	// 从上下文获取用户ID
	userID, ok := r.Context().Value("userID").(int64)
	if !ok {
		http.Error(w, "未登录", http.StatusUnauthorized)
		return
	}

	// 获取分页参数
	page := 1
	pageSize := 10

	if r.URL.Query().Get("page") != "" {
		p, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err == nil && p > 0 {
			page = p
		}
	}

	if r.URL.Query().Get("pageSize") != "" {
		ps, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
		if err == nil && ps > 0 {
			pageSize = ps
		}
	}

	// 获取支付记录
	payments, total, err := c.paymentService.GetPaymentsByUserID(userID, page, pageSize)
	if err != nil {
		http.Error(w, "获取支付记录失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 200,
		"msg":  "success",
		"data": map[string]interface{}{
			"list":     payments,
			"total":    total,
			"page":     page,
			"pageSize": pageSize,
		},
	})
}

// NotifyPayment 支付结果通知
func (c *PaymentController) NotifyPayment(w http.ResponseWriter, r *http.Request) {
	// 解析支付平台发送的通知
	// 注意：实际应用中需要验证通知的真实性

	var notifyData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&notifyData); err != nil {
		http.Error(w, "解析通知失败", http.StatusBadRequest)
		return
	}

	// 提取订单ID、交易ID和支付状态
	orderID, ok := notifyData["order_id"].(string)
	if !ok {
		http.Error(w, "缺少订单ID", http.StatusBadRequest)
		return
	}

	transactionID, ok := notifyData["transaction_id"].(string)
	if !ok {
		http.Error(w, "缺少交易ID", http.StatusBadRequest)
		return
	}

	status, ok := notifyData["status"].(string)
	if !ok {
		http.Error(w, "缺少支付状态", http.StatusBadRequest)
		return
	}

	// 更新支付状态
	if err := c.paymentService.UpdatePaymentStatus(orderID, transactionID, status); err != nil {
		http.Error(w, "更新支付状态失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 返回成功响应给支付平台
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 200,
		"msg":  "success",
	}) 
}