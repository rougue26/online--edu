package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"online-education-api/services"
)

// CommentController 评论控制器
 type CommentController struct {
	commentService services.CommentService
}

// NewCommentController 创建评论控制器实例
func NewCommentController(commentService services.CommentService) *CommentController {
	return &CommentController{
		commentService: commentService,
	}
}

// GetCommentList 获取评论列表
func (c *CommentController) GetCommentList(w http.ResponseWriter, r *http.Request) {
	// 获取查询参数
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("pageSize")
	videoIDStr := r.URL.Query().Get("videoID")

	// 默认值
	page := 1
	pageSize := 10
	videoID := 0

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

	if videoIDStr != "" {
		if vid, err := strconv.Atoi(videoIDStr); err == nil && vid > 0 {
			videoID = vid
		}
	}

	// 调用服务方法
	comments, total, err := c.commentService.GetCommentList(page, pageSize, videoID)
	if err != nil {
		http.Error(w, "获取评论列表失败: " + err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"list":  comments,
			"total": total,
		},
		"page":    page,
		"pageSize": pageSize,
	})
}

// DeleteComment 删除评论
func (c *CommentController) DeleteComment(w http.ResponseWriter, r *http.Request) {
	// 获取路径参数
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "无效的评论ID", http.StatusBadRequest)
		return
	}

	// 调用服务方法
	if err := c.commentService.DeleteComment(id); err != nil {
		http.Error(w, "删除评论失败: " + err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "评论删除成功",
	})
}