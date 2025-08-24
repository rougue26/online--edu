package services

import (
	"database/sql"
	"fmt"
	"log"

	"online-education-api/models"
)

// VideoService 视频服务接口
type VideoService interface {
	CreateVideo(video *models.Video) error
	GetVideoList(page, pageSize int, categoryID int) ([]models.Video, int, error)
	GetVideoByID(id int) (*models.Video, error)
	UpdateVideo(video *models.Video) error
	DeleteVideo(id int) error
	GetVideoCategories() ([]models.VideoCategory, error)
}

// videoService 视频服务实现
 type videoService struct {
	db *sql.DB
}

// NewVideoService 创建视频服务实例
func NewVideoService(db *sql.DB) VideoService {
	return &videoService{
	db: db,
	}
}

// CreateVideo 创建视频
func (s *videoService) CreateVideo(video *models.Video) error {
	query := `
	INSERT INTO videos (
		title, description, video_url, cover_image_url, duration, 
		author_id, category_id, is_public
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	result, err := s.db.Exec(
		query,
		video.Title,
		video.Description,
		video.VideoURL,
		video.CoverImageURL,
		video.Duration,
		video.AuthorID,
		video.CategoryID,
		video.IsPublic,
	)

	if err != nil {
		return fmt.Errorf("创建视频失败: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("获取视频ID失败: %w", err)
	}

	video.ID = int(id)
	return nil
}

// GetVideoList 获取视频列表
func (s *videoService) GetVideoList(page, pageSize int, categoryID int) ([]models.Video, int, error) {
	// 计算偏移量
	offset := (page - 1) * pageSize

	// 构建查询条件
	var query string
	var args []interface{}
	log.Println("获取视频列表，分类ID:", categoryID)

	if categoryID > 0 {
		query = `
		SELECT id, title, description, cover_image_url, duration, 
			created_at
		FROM videos
		WHERE category_id = ? 
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?
		`
		args = []interface{}{categoryID, pageSize, offset}
	} else {
		query = `
		SELECT id, title, description, cover_image_url, duration, 
			created_at
		FROM videos
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?
		`
		args = []interface{}{pageSize, offset}
	}

	// 查询视频列表
	log.Printf("执行SQL查询: %s", query)
	log.Printf("查询参数: %v", args)
	rows, err := s.db.Query(query, args...)
	if err != nil {
		log.Printf("SQL查询错误: %v", err)
		return nil, 0, fmt.Errorf("获取视频列表失败: %w", err)
	}
	defer rows.Close()

	var videos []models.Video
	for rows.Next() {
		var video models.Video
		err := rows.Scan(
			&video.ID,
			&video.Title,
			&video.Description,
			&video.CoverImageURL,
			&video.Duration,
			&video.CreatedAt,
		)
		if err != nil {
			log.Printf("扫描视频数据失败: %v", err)
			continue
		}
		videos = append(videos, video)
	}

	// 获取总记录数
	var total int
	if categoryID > 0 {
		err = s.db.QueryRow("SELECT COUNT(*) FROM videos WHERE category_id = ?", categoryID).Scan(&total)
	} else {
		err = s.db.QueryRow("SELECT COUNT(*) FROM videos").Scan(&total)
	}

	if err != nil {
		return videos, 0, fmt.Errorf("获取视频总数失败: %w", err)
	}

	return videos, total, nil
}

// GetVideoByID 获取视频详情
func (s *videoService) GetVideoByID(id int) (*models.Video, error) {
	query := `
	SELECT id, title, description, video_url, cover_image_url, duration, 
			author_id, category_id, view_count, like_count, favorite_count, 
			is_public, created_at, updated_at
	FROM videos
	WHERE id = ?
	`

	var video models.Video
	err := s.db.QueryRow(query, id).Scan(
		&video.ID,
		&video.Title,
		&video.Description,
		&video.VideoURL,
		&video.CoverImageURL,
		&video.Duration,
		&video.AuthorID,
		&video.CategoryID,
		&video.ViewCount,
		&video.LikeCount,
		&video.FavoriteCount,
		&video.IsPublic,
		&video.CreatedAt,
		&video.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("视频不存在")
		}
		return nil, fmt.Errorf("获取视频详情失败: %w", err)
	}

	// 更新观看次数
	_, err = s.db.Exec("UPDATE videos SET view_count = view_count + 1 WHERE id = ?", id)
	if err != nil {
		log.Printf("更新视频观看次数失败: %v", err)
	}

	return &video, nil
}

// UpdateVideo 更新视频
func (s *videoService) UpdateVideo(video *models.Video) error {
	query := `
	UPDATE videos SET
		title = ?, description = ?, cover_image_url = ?, 
		category_id = ?, is_public = ?
	WHERE id = ? AND author_id = ?
	`

	result, err := s.db.Exec(
		query,
		video.Title,
		video.Description,
		video.CoverImageURL,
		video.CategoryID,
		video.IsPublic,
		video.ID,
		video.AuthorID,
	)

	if err != nil {
		return fmt.Errorf("更新视频失败: %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("获取受影响行数失败: %w", err)
	}

	if affected == 0 {
		return fmt.Errorf("视频不存在或无权限更新")
	}

	return nil
}

// DeleteVideo 删除视频
func (s *videoService) DeleteVideo(id int) error {
	query := `
	DELETE FROM videos WHERE id = ?
	`

	result, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("删除视频失败: %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("获取受影响行数失败: %w", err)
	}

	if affected == 0 {
		return fmt.Errorf("视频不存在")
	}

	return nil
}

// GetVideoCategories 获取视频分类
func (s *videoService) GetVideoCategories() ([]models.VideoCategory, error) {
	query := `
	SELECT id, name, description, created_at
	FROM video_categories
	ORDER BY created_at ASC
	`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("获取视频分类失败: %w", err)
	}
	defer rows.Close()

	var categories []models.VideoCategory
	for rows.Next() {
		var category models.VideoCategory
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&category.CreatedAt,
		)
		if err != nil {
			log.Printf("扫描分类数据失败: %v", err)
			continue
		}
		categories = append(categories, category)
	}

	return categories, nil
}