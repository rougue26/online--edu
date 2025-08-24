package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"online-education-api/services"
)

// UserCourseController 用户课程控制器
type UserCourseController struct {
	userCourseService services.UserCourseService
}

// NewUserCourseController 创建用户课程控制器实例
func NewUserCourseController(userCourseService services.UserCourseService) *UserCourseController {
	return &UserCourseController{
		userCourseService: userCourseService,
	}
}

// EnrollCourse 用户报名课程
func (c *UserCourseController) EnrollCourse(w http.ResponseWriter, r *http.Request) {
	// 从上下文获取用户ID
	userID, ok := r.Context().Value("userID").(int64)
	if !ok {
		http.Error(w, "未登录", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	courseIDStr := vars["courseID"]

	// 转换课程ID为int64
	courseID, err := strconv.ParseInt(courseIDStr, 10, 64)
	if err != nil {
		http.Error(w, "无效的课程ID", http.StatusBadRequest)
		return
	}

	if err := c.userCourseService.EnrollCourse(userID, courseID); err != nil {
		http.Error(w, "报名失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 200,
		"msg":  "报名成功",
	})
}

// GetUserCourses 获取用户报名的课程列表
func (c *UserCourseController) GetUserCourses(w http.ResponseWriter, r *http.Request) {
	// 从上下文获取用户ID
	userID, ok := r.Context().Value("userID").(int64)
	if !ok {
		http.Error(w, "未登录", http.StatusUnauthorized)
		return
	}

	// 获取查询参数
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

	courses, total, err := c.userCourseService.GetUserCourses(userID, page, pageSize)
	if err != nil {
		http.Error(w, "获取课程列表失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 200,
		"msg":  "success",
		"data": map[string]interface{}{
			"list":  courses,
			"total": total,
			"page":  page,
			"pageSize": pageSize,
		},
	})
}

// GetUserCourseByID 检查用户是否报名了某课程
func (c *UserCourseController) GetUserCourseByID(w http.ResponseWriter, r *http.Request) {
	// 从上下文获取用户ID
	userID, ok := r.Context().Value("userID").(int64)
	if !ok {
		http.Error(w, "未登录", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	courseIDStr := vars["courseID"]

	// 转换课程ID为int64
	courseID, err := strconv.ParseInt(courseIDStr, 10, 64)
	if err != nil {
		http.Error(w, "无效的课程ID", http.StatusBadRequest)
		return
	}

	userCourse, err := c.userCourseService.GetUserCourseByID(userID, courseID)
	if err != nil {
		http.Error(w, "获取报名记录失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 200,
		"msg":  "success",
		"data": userCourse,
	})
}

// UnenrollCourse 用户取消报名课程
func (c *UserCourseController) UnenrollCourse(w http.ResponseWriter, r *http.Request) {
	// 从上下文获取用户ID
	userID, ok := r.Context().Value("userID").(int64)
	if !ok {
		http.Error(w, "未登录", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	courseIDStr := vars["courseID"]

	// 转换课程ID为int64
	courseID, err := strconv.ParseInt(courseIDStr, 10, 64)
	if err != nil {
		http.Error(w, "无效的课程ID", http.StatusBadRequest)
		return
	}

	if err := c.userCourseService.UnenrollCourse(userID, courseID); err != nil {
		http.Error(w, "取消报名失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 200,
		"msg":  "取消报名成功",
	})
}