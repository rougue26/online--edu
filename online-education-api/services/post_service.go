package services

import (
	"database/sql"
	"errors"
	"online-education-api/models"
	"time"
)

// PostService 帖子服务接口
type PostService interface {
	GetPostList(page, pageSize int, search string) ([]*models.PostResponse, int, error)
	GetPostDetail(id int64) (*models.PostResponse, error)
	CreatePost(post *models.Post) error
	UpdatePost(post *models.Post) error
	DeletePost(id int64) error
	GetUserPosts(userID int64, page, pageSize int) ([]*models.PostResponse, int, error)
}

// postService 帖子服务实现
 type postService struct {
	db *sql.DB
}

// NewPostService 创建帖子服务实例
func NewPostService(db *sql.DB) PostService {
	return &postService{db: db}
}

// GetPostList 获取帖子列表
func (s *postService) GetPostList(page, pageSize int, search string) ([]*models.PostResponse, int, error) {
	// 计算偏移量
	offset := (page - 1) * pageSize

	// 构建查询条件
	query := `SELECT p.id, p.user_id, p.title, p.content, p.category, p.views, p.comment_count, p.like_count, CASE p.status WHEN 'published' THEN 1 WHEN 'draft' THEN 0 ELSE 0 END as status, p.created_at, p.updated_at, u.username, u.avatar FROM posts p LEFT JOIN users u ON p.user_id = u.id WHERE p.status IN ('published', 1)`
countQuery := `SELECT COUNT(*) FROM posts WHERE status IN ('published', 1)`

	// 添加查询参数
	params := []interface{}{}

	if search != "" {
		query += ` AND (p.title LIKE ? OR p.content LIKE ?)`
		countQuery += ` AND (title LIKE ? OR content LIKE ?)`
		searchParam := "%" + search + "%"
		params = append(params, searchParam, searchParam)
	}

	// 添加排序和分页 - 确保最新创建的帖子在前面
	query += ` ORDER BY p.created_at DESC, p.id DESC LIMIT ? OFFSET ?`
	params = append(params, pageSize, offset)

	// 执行查询
	rows, err := s.db.Query(query, params...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// 获取总记录数
	var total int
	var countParams []interface{}
	if len(params) > 2 {
		countParams = params[:len(params)-2] // 移除limit和offset参数
	} else {
		countParams = []interface{}{}
	}
	if err := s.db.QueryRow(countQuery, countParams...).Scan(&total); err != nil {
		return nil, 0, err
	}

	// 解析结果
	var posts []*models.PostResponse
	for rows.Next() {
		var post models.PostResponse
		var avatar sql.NullString
			
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.Category, &post.ViewCount, &post.CommentCount, &post.LikeCount, &post.Status, &post.CreatedAt, &post.UpdatedAt, &post.UserName, &avatar); err != nil {
			return nil, 0, err
		}
			
		if avatar.Valid {
			post.Avatar = avatar.String
		}
			
		posts = append(posts, &post)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

// GetPostDetail 获取帖子详情
func (s *postService) GetPostDetail(id int64) (*models.PostResponse, error) {
	// 查询帖子信息
	query := `SELECT p.id, p.user_id, p.title, p.content, p.category, p.views, p.comment_count, p.like_count, CASE p.status WHEN 'published' THEN 1 WHEN 'draft' THEN 0 ELSE 0 END as status, p.created_at, p.updated_at, u.username, u.avatar FROM posts p LEFT JOIN users u ON p.user_id = u.id WHERE p.id = ?`
	row := s.db.QueryRow(query, id)

	var post models.PostResponse
	var avatar sql.NullString
		
	if err := row.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.Category, &post.ViewCount, &post.CommentCount, &post.LikeCount, &post.Status, &post.CreatedAt, &post.UpdatedAt, &post.UserName, &avatar); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("帖子不存在")
		}
		return nil, err
	}

	if avatar.Valid {
		post.Avatar = avatar.String
	}

	// 更新浏览量
	query = `UPDATE posts SET views = views + 1 WHERE id = ?`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return nil, err
	}

	post.ViewCount++

	return &post, nil
}

// CreatePost 创建帖子
func (s *postService) CreatePost(post *models.Post) error {
	query := `INSERT INTO posts (user_id, title, content, category, views, comment_count, like_count, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	now := time.Now()

	// 默认值
	post.ViewCount = 0
	post.CommentCount = 0
	post.LikeCount = 0
	post.Status = 1 // 1: published
	if post.Category == "" {
		post.Category = "未分类"
	}

	// 将int状态转换为字符串
	statusStr := "draft"
	if post.Status == 1 {
		statusStr = "published"
	}

	result, err := s.db.Exec(query, post.UserID, post.Title, post.Content, post.Category, post.ViewCount, post.CommentCount, post.LikeCount, statusStr, now, now)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	post.ID = id

	return nil
}

// UpdatePost 更新帖子
func (s *postService) UpdatePost(post *models.Post) error {
	query := `UPDATE posts SET title = ?, content = ?, category = ?, status = ?, updated_at = ? WHERE id = ?`
	now := time.Now()

	// 将int状态转换为字符串
	statusStr := "draft"
	if post.Status == 1 {
		statusStr = "published"
	} else if post.Status == 2 {
		statusStr = "archived"
	}

	result, err := s.db.Exec(query, post.Title, post.Content, post.Category, statusStr, now, post.ID)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("帖子不存在")
	}

	return nil
}

// DeletePost 删除帖子
func (s *postService) DeletePost(id int64) error {
	// 软删除，更新状态为归档
	query := `UPDATE posts SET status = "archived", updated_at = ? WHERE id = ?`
	now := time.Now()

	result, err := s.db.Exec(query, now, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("帖子不存在")
	}

	return nil
}

// GetUserPosts 获取用户发布的帖子
func (s *postService) GetUserPosts(userID int64, page, pageSize int) ([]*models.PostResponse, int, error) {
	// 计算偏移量
	offset := (page - 1) * pageSize

	// 查询用户帖子
	query := `SELECT p.id, p.user_id, p.title, p.content, p.views, p.comment_count, p.like_count, p.status, p.created_at, p.updated_at, u.username, u.avatar FROM posts p LEFT JOIN users u ON p.user_id = u.id WHERE p.user_id = ? ORDER BY p.created_at DESC LIMIT ? OFFSET ?`
	rows, err := s.db.Query(query, userID, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// 获取总记录数
	var total int
	countQuery := `SELECT COUNT(*) FROM posts WHERE user_id = ?`
	if err := s.db.QueryRow(countQuery, userID).Scan(&total); err != nil {
		return nil, 0, err
	}

	// 解析结果
	var posts []*models.PostResponse
	for rows.Next() {
		var post models.PostResponse
		var avatar sql.NullString
		
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.ViewCount, &post.CommentCount, &post.LikeCount, &post.Status, &post.CreatedAt, &post.UpdatedAt, &post.UserName, &avatar); err != nil {
			return nil, 0, err
		}
		
		if avatar.Valid {
			post.Avatar = avatar.String
		}
		
		posts = append(posts, &post)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}