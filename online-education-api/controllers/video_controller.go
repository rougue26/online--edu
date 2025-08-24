package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"online-education-api/models"
	"online-education-api/services"
)

// VideoController 视频控制器
type VideoController struct {
	videoService services.VideoService
}

// NewVideoController 创建视频控制器实例
func NewVideoController(videoService services.VideoService) *VideoController {
	return &VideoController{
		videoService: videoService,
	}
}

// CreateVideo 创建视频
func (c *VideoController) CreateVideo(w http.ResponseWriter, r *http.Request) {
	var video models.Video
	if err := json.NewDecoder(r.Body).Decode(&video); err != nil {
		http.Error(w, "无效的请求体: " + err.Error(), http.StatusBadRequest)
		return
	}

	// 在实际应用中，应该从认证信息中获取用户ID
	// video.AuthorID = getUserIDFromContext(r.Context())

	if err := c.videoService.CreateVideo(&video); err != nil {
		http.Error(w, "创建视频失败: " + err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    video,
	})
}

// GetVideoList 获取视频列表
func (c *VideoController) GetVideoList(w http.ResponseWriter, r *http.Request) {
	// 获取查询参数
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("pageSize")
	categoryIDStr := r.URL.Query().Get("categoryID")

	// 默认值
	page := 1
	pageSize := 10
	categoryID := 0

	// 解析参数
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

	if categoryIDStr != "" {
		if cid, err := strconv.Atoi(categoryIDStr); err == nil && cid > 0 {
			categoryID = cid
		}
	}

	// 调用服务方法
	videos, total, err := c.videoService.GetVideoList(page, pageSize, categoryID)
	if err != nil {
		http.Error(w, "获取视频列表失败: " + err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    videos,
		"total":   total,
		"page":    page,
		"pageSize": pageSize,
	})
}

// GetVideoByID 获取视频详情
func (c *VideoController) GetVideoByID(w http.ResponseWriter, r *http.Request) {
	// 获取路径参数
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "无效的视频ID", http.StatusBadRequest)
		return
	}

	// 调用服务方法
	video, err := c.videoService.GetVideoByID(id)
	if err != nil {
		http.Error(w, "获取视频详情失败: " + err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    video,
	})
}

// UpdateVideo 更新视频
func (c *VideoController) UpdateVideo(w http.ResponseWriter, r *http.Request) {
	// 获取路径参数
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "无效的视频ID", http.StatusBadRequest)
		return
	}

	// 解析请求体
	var video models.Video
	if err := json.NewDecoder(r.Body).Decode(&video); err != nil {
		http.Error(w, "无效的请求体: " + err.Error(), http.StatusBadRequest)
		return
	}

	// 设置视频ID
	video.ID = id

	// 在实际应用中，应该从认证信息中获取用户ID并验证权限
	// video.AuthorID = getUserIDFromContext(r.Context())

	// 调用服务方法
	if err := c.videoService.UpdateVideo(&video); err != nil {
		http.Error(w, "更新视频失败: " + err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "视频更新成功",
	})
}

// DeleteVideo 删除视频
func (c *VideoController) DeleteVideo(w http.ResponseWriter, r *http.Request) {
	// 获取路径参数
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "无效的视频ID", http.StatusBadRequest)
		return
	}

	// 在实际应用中，应该验证用户权限

	// 调用服务方法
	if err := c.videoService.DeleteVideo(id); err != nil {
		http.Error(w, "删除视频失败: " + err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "视频删除成功",
	})
}

// GetVideoCategories 获取视频分类
func (c *VideoController) GetVideoCategories(w http.ResponseWriter, r *http.Request) {
	// 调用服务方法
	categories, err := c.videoService.GetVideoCategories()
	if err != nil {
		http.Error(w, "获取视频分类失败: " + err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    categories,
	})
}