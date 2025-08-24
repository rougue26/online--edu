package services

import (
	"database/sql"
	"fmt"
	"log"

	"online-education-api/models"
)

// CommentService 评论服务接口
type CommentService interface {
	GetCommentList(page, pageSize int, videoID int) ([]models.VideoCommentWithUserInfo, int, error)
	DeleteComment(id int) error
}

// commentService 评论服务实现
type commentService struct {
	db *sql.DB
}

// NewCommentService 创建评论服务实例
func NewCommentService(db *sql.DB) CommentService {
	return &commentService{
	db: db,
	}
}

// GetCommentList 获取评论列表
func (s *commentService) GetCommentList(page, pageSize int, videoID int) ([]models.VideoCommentWithUserInfo, int, error) {
	// 计算偏移量
	offset := (page - 1) * pageSize

	// 构建查询条件
	var query string
	var args []interface{}

	if videoID > 0 {
		query = `
		SELECT vc.id, vc.video_id, vc.user_id, vc.content, vc.created_at, vc.updated_at,
		       u.username as author_name, v.title as video_title
		FROM video_comments vc
		JOIN users u ON vc.user_id = u.id
		JOIN videos v ON vc.video_id = v.id
		WHERE vc.video_id = ?
		ORDER BY vc.created_at DESC
		LIMIT ? OFFSET ?
		`
		args = []interface{}{videoID, pageSize, offset}
	} else {
		query = `
		SELECT vc.id, vc.video_id, vc.user_id, vc.content, vc.created_at, vc.updated_at,
		       u.username as author_name, v.title as video_title
		FROM video_comments vc
		JOIN users u ON vc.user_id = u.id
		JOIN videos v ON vc.video_id = v.id
		ORDER BY vc.created_at DESC
		LIMIT ? OFFSET ?
		`
		args = []interface{}{pageSize, offset}
	}

	// 查询评论列表
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("获取评论列表失败: %w", err)
	}
	defer rows.Close()

	var comments []models.VideoCommentWithUserInfo
	for rows.Next() {
		var comment models.VideoCommentWithUserInfo
		err := rows.Scan(
			&comment.ID,
			&comment.VideoID,
			&comment.UserID,
			&comment.Content,
			&comment.CreatedAt,
			&comment.UpdatedAt,
			&comment.AuthorName,
			&comment.VideoTitle,
		)
		if err != nil {
			log.Printf("扫描评论数据失败: %v", err)
			continue
		}
		comments = append(comments, comment)
	}

	// 获取总记录数
	var total int
	if videoID > 0 {
		err = s.db.QueryRow("SELECT COUNT(*) FROM video_comments WHERE video_id = ?", videoID).Scan(&total)
	} else {
		err = s.db.QueryRow("SELECT COUNT(*) FROM video_comments").Scan(&total)
	}

	if err != nil {
		return comments, 0, fmt.Errorf("获取评论总数失败: %w", err)
	}

	return comments, total, nil
}

// DeleteComment 删除评论
func (s *commentService) DeleteComment(id int) error {
	query := `
	DELETE FROM video_comments WHERE id = ?
	`

	result, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("删除评论失败: %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("获取受影响行数失败: %w", err)
	}

	if affected == 0 {
		return fmt.Errorf("评论不存在")
	}

	return nil
}