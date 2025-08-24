package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"online-education-api/models"
	"online-education-api/services"
)

// PostController 帖子控制器
type PostController struct {
	postService services.PostService
}

// NewPostController 创建帖子控制器实例
func NewPostController(postService services.PostService) *PostController {
	return &PostController{
		postService: postService,
	}
}

// GetPostList 获取帖子列表
func (c *PostController) GetPostList(w http.ResponseWriter, r *http.Request) {
	// 获取查询参数
	page := 1
	pageSize := 10
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

	if r.URL.Query().Get("search") != "" {
		search = r.URL.Query().Get("search")
	}

	posts, total, err := c.postService.GetPostList(page, pageSize, search)
	if err != nil {
		http.Error(w, "获取帖子列表失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 200,
		"msg":  "success",
		"data": map[string]interface{}{
			"list":  posts,
			"total": total,
			"page":  page,
			"pageSize": pageSize,
		},
	})
}

// GetPostDetail 获取帖子详情
func (c *PostController) GetPostDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// 转换ID为int64
	postID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "无效的帖子ID", http.StatusBadRequest)
		return
	}

	post, err := c.postService.GetPostDetail(postID)
	if err != nil {
		http.Error(w, "获取帖子详情失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 200,
		"msg":  "success",
		"data": post,
	})
}

// CreatePost 创建帖子
func (c *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
	// 从上下文获取用户ID
	userID, ok := r.Context().Value("userID").(int64)
	if !ok {
		http.Error(w, "未登录", http.StatusUnauthorized)
		return
	}

	var req models.CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "请求参数错误", http.StatusBadRequest)
		return
	}

	// 验证参数
	// 这里应该添加参数验证逻辑

	// 创建帖子
	post := &models.Post{
		UserID:  userID,
		Title:   req.Title,
		Content: req.Content,
		Category: req.Category,
	}

	if err := c.postService.CreatePost(post); err != nil {
		http.Error(w, "创建帖子失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 200,
		"msg":  "创建成功",
		"data": post,
	})
}

// UpdatePost 更新帖子
func (c *PostController) UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// 转换ID为int64
	postID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "无效的帖子ID", http.StatusBadRequest)
		return
	}

	// 从上下文获取用户ID
	userID, ok := r.Context().Value("userID").(int64)
	if !ok {
		http.Error(w, "未登录", http.StatusUnauthorized)
		return
	}

	// 检查帖子是否存在，并且是否为当前用户创建
	post, err := c.postService.GetPostDetail(postID)
	if err != nil {
		http.Error(w, "获取帖子失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if post.UserID != userID {
		http.Error(w, "无权限修改该帖子", http.StatusForbidden)
		return
	}

	var req models.UpdatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "请求参数错误", http.StatusBadRequest)
		return
	}

	// 更新帖子信息
	if req.Title != "" {
		post.Title = req.Title
	}
	if req.Content != "" {
		post.Content = req.Content
	}
	if req.Status >= 0 {
		post.Status = req.Status
	}

	// 转换为models.Post类型
	updatePost := &models.Post{
		ID:      post.ID,
		UserID:  post.UserID,
		Title:   post.Title,
		Content: post.Content,
		Status:  post.Status,
	}

	if err := c.postService.UpdatePost(updatePost); err != nil {
		http.Error(w, "更新帖子失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 200,
		"msg":  "更新成功",
		"data": updatePost,
	})
}

// DeletePost 删除帖子
func (c *PostController) DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// 转换ID为int64
	postID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "无效的帖子ID", http.StatusBadRequest)
		return
	}

	// 从上下文获取用户ID
	userID, ok := r.Context().Value("userID").(int64)
	if !ok {
		http.Error(w, "未登录", http.StatusUnauthorized)
		return
	}

	// 检查帖子是否存在，并且是否为当前用户创建
	post, err := c.postService.GetPostDetail(postID)
	if err != nil {
		http.Error(w, "获取帖子失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if post.UserID != userID {
		http.Error(w, "无权限删除该帖子", http.StatusForbidden)
		return
	}

	if err := c.postService.DeletePost(postID); err != nil {
		http.Error(w, "删除帖子失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 200,
		"msg":  "删除成功",
	})
}

// GetUserPosts 获取用户发布的帖子
func (c *PostController) GetUserPosts(w http.ResponseWriter, r *http.Request) {
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

	posts, total, err := c.postService.GetUserPosts(userID, page, pageSize)
	if err != nil {
		http.Error(w, "获取帖子列表失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 200,
		"msg":  "success",
		"data": map[string]interface{}{
			"list":  posts,
			"total": total,
			"page":  page,
			"pageSize": pageSize,
		},
	})
}