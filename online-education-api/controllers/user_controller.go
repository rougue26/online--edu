package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"online-education-api/models"
	"online-education-api/services"
)

// UserController 用户控制器
type UserController struct {
	userService services.UserService
}

// NewUserController 创建用户控制器实例
func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

// Register 处理用户注册请求
func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var registerReq models.UserRegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&registerReq); err != nil {
		http.Error(w, "无效的请求数据", http.StatusBadRequest)
		return
	}

	user, err := c.userService.Register(&registerReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 转换为响应模型
	response := models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Nickname:  user.Nickname,
		Bio:       user.Bio,
		CreatedAt: user.CreatedAt,
		Status:    user.Status,
		Role:      user.Role,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// Login 处理用户登录请求
func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq models.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, "无效的请求数据", http.StatusBadRequest)
		return
	}

	user, token, err := c.userService.Login(&loginReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// 转换为响应模型
	response := models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Nickname:  user.Nickname,
		Bio:       user.Bio,
		CreatedAt: user.CreatedAt,
		Status:    user.Status,
		Role:      user.Role,
	}

	// 返回用户信息和令牌
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user":  response,
		"token": token,
	})
}

// GetProfile 获取用户个人资料
func (c *UserController) GetProfile(w http.ResponseWriter, r *http.Request) {
	// 从请求上下文中获取用户ID
	userID, ok := r.Context().Value("userID").(int64)
	if !ok {
		http.Error(w, "无法获取用户信息", http.StatusUnauthorized)
		return
	}

	user, err := c.userService.GetUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// 转换为响应模型
	response := models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Nickname:  user.Nickname,
		Bio:       user.Bio,
		CreatedAt: user.CreatedAt,
		Status:    user.Status,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UpdateProfile 更新用户个人资料
func (c *UserController) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	// 从请求上下文中获取用户ID
	userID, ok := r.Context().Value("userID").(int64)
	if !ok {
		http.Error(w, "无法获取用户信息", http.StatusUnauthorized)
		return
	}

	var updateReq models.UserUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
		http.Error(w, "无效的请求数据", http.StatusBadRequest)
		return
	}

	user, err := c.userService.UpdateUser(userID, &updateReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 转换为响应模型
	response := models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Nickname:  user.Nickname,
		Bio:       user.Bio,
		CreatedAt: user.CreatedAt,
		Status:    user.Status,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// ChangePassword 处理修改密码请求
func (c *UserController) ChangePassword(w http.ResponseWriter, r *http.Request) {
	// 从请求上下文中获取用户ID
	userID, ok := r.Context().Value("userID").(int64)
	if !ok {
		http.Error(w, "无法获取用户信息", http.StatusUnauthorized)
		return
	}

	var req struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "无效的请求数据", http.StatusBadRequest)
		return
	}

	err := c.userService.ChangePassword(userID, req.OldPassword, req.NewPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "密码修改成功"})
}

// GetUserList 获取用户列表
func (c *UserController) GetUserList(w http.ResponseWriter, r *http.Request) {
	// 检查用户角色
	role, ok := r.Context().Value("role").(string)
	if !ok || role != "admin" {
		http.Error(w, "权限不足，需要管理员角色", http.StatusForbidden)
		return
	}
	
	// 获取分页参数
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("pageSize")
	
	// 默认为第1页，每页10条
	page := 1
	pageSize := 10
	
	// 转换分页参数
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}
	
	if pageSizeStr != "" {
		if ps, err := strconv.Atoi(pageSizeStr); err == nil && ps > 0 {
			pageSize = ps
		}
	}
	
	// 调用服务层获取用户列表
	users, total, err := c.userService.GetUserList(page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// 转换为响应模型
	var userResponses []models.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, models.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Avatar:    user.Avatar,
			Nickname:  user.Nickname,
			Bio:       user.Bio,
			CreatedAt: user.CreatedAt,
			Status:    user.Status,
			Role:      user.Role,
		})
	}
	
	// 返回用户列表和总数
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"list":  userResponses,
		"total": total,
	})
}

// CreateUser 创建新用户
func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	// 检查用户角色
	role, ok := r.Context().Value("role").(string)
	if !ok || role != "admin" {
		http.Error(w, "权限不足，需要管理员角色", http.StatusForbidden)
		return
	}
	
	var createReq models.UserCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&createReq); err != nil {
		http.Error(w, "无效的请求数据", http.StatusBadRequest)
		return
	}

	user, err := c.userService.CreateUser(&createReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 转换为响应模型
	response := models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Nickname:  user.Nickname,
		Bio:       user.Bio,
		CreatedAt: user.CreatedAt,
		Status:    user.Status,
		Role:      user.Role,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetUserByID 根据ID获取用户详情
func (c *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// 从URL路径中获取用户ID
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "无效的用户ID", http.StatusBadRequest)
		return
	}

	user, err := c.userService.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// 转换为响应模型
	response := models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Nickname:  user.Nickname,
		Bio:       user.Bio,
		CreatedAt: user.CreatedAt,
		Status:    user.Status,
		Role:      user.Role,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UpdateUser 更新用户信息
func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// 检查用户角色
	role, ok := r.Context().Value("role").(string)
	if !ok || role != "admin" {
		http.Error(w, "权限不足，需要管理员角色", http.StatusForbidden)
		return
	}
	
	// 从URL路径中获取用户ID
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "无效的用户ID", http.StatusBadRequest)
		return
	}

	var updateReq models.UserUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
		http.Error(w, "无效的请求数据", http.StatusBadRequest)
		return
	}

	user, err := c.userService.UpdateUser(id, &updateReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 转换为响应模型
	response := models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Nickname:  user.Nickname,
		Bio:       user.Bio,
		CreatedAt: user.CreatedAt,
		Status:    user.Status,
		Role:      user.Role,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteUser 删除用户
func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// 检查用户角色
	role, ok := r.Context().Value("role").(string)
	if !ok || role != "admin" {
		http.Error(w, "权限不足，需要管理员角色", http.StatusForbidden)
		return
	}
	
	// 从URL路径中获取用户ID
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "无效的用户ID", http.StatusBadRequest)
		return
	}

	err = c.userService.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "用户删除成功"})
}