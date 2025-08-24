package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"online-education-api/models"
	"online-education-api/services"
)

// CourseCategoryController 课程分类控制器
type CourseCategoryController struct {
	courseCategoryService services.CourseCategoryService
}

// NewCourseCategoryController 创建课程分类控制器实例
func NewCourseCategoryController(courseCategoryService services.CourseCategoryService) *CourseCategoryController {
	return &CourseCategoryController{
		courseCategoryService: courseCategoryService,
	}
}

// GetAllCategories 获取所有课程分类
func (c *CourseCategoryController) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := c.courseCategoryService.GetAllCategories()
	if err != nil {
		http.Error(w, "获取分类失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 200,
		"msg":  "success",
		"data": categories,
	})
}

// GetCategoryByID 根据ID获取课程分类
func (c *CourseCategoryController) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// 转换ID为int64
	categoryID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "无效的分类ID", http.StatusBadRequest)
		return
	}

	category, err := c.courseCategoryService.GetCategoryByID(categoryID)
	if err != nil {
		http.Error(w, "获取分类失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 200,
		"msg":  "success",
		"data": category,
	})
}

// CreateCategory 创建课程分类
func (c *CourseCategoryController) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.CourseCategory
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, "请求参数错误", http.StatusBadRequest)
		return
	}

	if err := c.courseCategoryService.CreateCategory(&category); err != nil {
		http.Error(w, "创建分类失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 200,
		"msg":  "创建成功",
		"data": category,
	})
}

// UpdateCategory 更新课程分类
func (c *CourseCategoryController) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// 转换ID为int64
	categoryID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "无效的分类ID", http.StatusBadRequest)
		return
	}

	var category models.CourseCategory
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, "请求参数错误", http.StatusBadRequest)
		return
	}

	category.ID = categoryID
	if err := c.courseCategoryService.UpdateCategory(&category); err != nil {
		http.Error(w, "更新分类失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 200,
		"msg":  "更新成功",
		"data": category,
	})
}

// DeleteCategory 删除课程分类
func (c *CourseCategoryController) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// 转换ID为int64
	categoryID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "无效的分类ID", http.StatusBadRequest)
		return
	}

	if err := c.courseCategoryService.DeleteCategory(categoryID); err != nil {
		http.Error(w, "删除分类失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 200,
		"msg":  "删除成功",
	})
}