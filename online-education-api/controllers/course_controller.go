package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"online-education-api/models"
	"online-education-api/services"
)

// CourseController 课程控制器
type CourseController struct {
	courseService services.CourseService
}

// NewCourseController 创建课程控制器实例
func NewCourseController(courseService services.CourseService) *CourseController {
	return &CourseController{
		courseService: courseService,
	}
}

// GetCourseList 获取课程列表
func (c *CourseController) GetCourseList(w http.ResponseWriter, r *http.Request) {
	// 获取查询参数
	page := 1
	pageSize := 10
	categoryID := int64(0)
	level := int64(0)
	search := ""

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

	if r.URL.Query().Get("categoryID") != "" {
		cid, err := strconv.ParseInt(r.URL.Query().Get("categoryID"), 10, 64)
		if err == nil && cid > 0 {
			categoryID = cid
		}
	}

	if r.URL.Query().Get("level") != "" {
		l, err := strconv.ParseInt(r.URL.Query().Get("level"), 10, 64)
		if err == nil && l > 0 {
			level = l
		}
	}

	if r.URL.Query().Get("search") != "" {
		search = r.URL.Query().Get("search")
	}

	courses, total, err := c.courseService.GetCourseList(page, pageSize, categoryID, level, search)
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

// GetCourseDetail 获取课程详情
func (c *CourseController) GetCourseDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// 转换ID为int64
	courseID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "无效的课程ID", http.StatusBadRequest)
		return
	}

	course, err := c.courseService.GetCourseDetail(courseID)
	if err != nil {
		http.Error(w, "获取课程详情失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 200,
		"msg":  "success",
		"data": course,
	})
}

// CreateCourse 创建课程
func (c *CourseController) CreateCourse(w http.ResponseWriter, r *http.Request) {
	// 从上下文获取用户ID
	userID, ok := r.Context().Value("userID").(int64)
	if !ok {
		http.Error(w, "未登录", http.StatusUnauthorized)
		return
	}

	var req models.CreateCourseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "请求参数错误", http.StatusBadRequest)
		return
	}

	// 验证参数
	// 这里应该添加参数验证逻辑

	// 创建课程
	course := &models.Course{
		Title:         req.Title,
		Description:   req.Description,
		CoverImage:    req.CoverImage,
		Price:         req.Price,
		OriginalPrice: req.OriginalPrice,
		CategoryID:    req.CategoryID,
		TeacherID:     userID,
		Level:         req.Level,
		Status:        req.Status,
	}

	if err := c.courseService.CreateCourse(course); err != nil {
		http.Error(w, "创建课程失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 200,
		"msg":  "创建成功",
		"data": course,
	})
}

// UpdateCourse 更新课程
func (c *CourseController) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// 转换ID为int64
	courseID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "无效的课程ID", http.StatusBadRequest)
		return
	}

	// 从上下文获取用户ID
	userID, ok := r.Context().Value("userID").(int64)
	if !ok {
		http.Error(w, "未登录", http.StatusUnauthorized)
		return
	}

	// 检查课程是否存在，并且是否为当前用户创建
	course, err := c.courseService.GetCourseDetail(courseID)
	if err != nil {
		http.Error(w, "获取课程失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if course.TeacherID != userID {
		http.Error(w, "无权限修改该课程", http.StatusForbidden)
		return
	}

	var req models.UpdateCourseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "请求参数错误", http.StatusBadRequest)
		return
	}

	// 更新课程信息
	if req.Title != "" {
		course.Title = req.Title
	}
	if req.Description != "" {
		course.Description = req.Description
	}
	if req.CoverImage != "" {
		course.CoverImage = req.CoverImage
	}
	if req.Price > 0 {
		course.Price = req.Price
	}
	if req.OriginalPrice > 0 {
		course.OriginalPrice = req.OriginalPrice
	}
	if req.CategoryID > 0 {
		course.CategoryID = req.CategoryID
	}
	if req.Level > 0 {
		course.Level = req.Level
	}
	if req.Status >= 0 {
		course.Status = req.Status
	}

	if err := c.courseService.UpdateCourse(&course.Course); err != nil {
		http.Error(w, "更新课程失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 200,
		"msg":  "更新成功",
		"data": course,
	})
}

// DeleteCourse 删除课程
func (c *CourseController) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// 转换ID为int64
	courseID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "无效的课程ID", http.StatusBadRequest)
		return
	}

	// 从上下文获取用户ID
	userID, ok := r.Context().Value("userID").(int64)
	if !ok {
		http.Error(w, "未登录", http.StatusUnauthorized)
		return
	}

	// 检查课程是否存在，并且是否为当前用户创建
	course, err := c.courseService.GetCourseDetail(courseID)
	if err != nil {
		http.Error(w, "获取课程失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if course.TeacherID != userID {
		http.Error(w, "无权限删除该课程", http.StatusForbidden)
		return
	}

	if err := c.courseService.DeleteCourse(courseID); err != nil {
		http.Error(w, "删除课程失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 200,
		"msg":  "删除成功",
	})
}