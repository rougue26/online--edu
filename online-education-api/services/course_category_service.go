package services

import (
	"database/sql"
	"errors"
	"online-education-api/models"
	"time"
)

// CourseCategoryService 课程分类服务接口
type CourseCategoryService interface {
	GetAllCategories() ([]*models.CourseCategoryResponse, error)
	GetCategoryByID(id int64) (*models.CourseCategoryResponse, error)
	CreateCategory(category *models.CourseCategory) error
	UpdateCategory(category *models.CourseCategory) error
	DeleteCategory(id int64) error
}

// courseCategoryService 课程分类服务实现
 type courseCategoryService struct {
	db *sql.DB
}

// NewCourseCategoryService 创建课程分类服务实例
func NewCourseCategoryService(db *sql.DB) CourseCategoryService {
	return &courseCategoryService{db: db}
}

// GetAllCategories 获取所有课程分类
func (s *courseCategoryService) GetAllCategories() ([]*models.CourseCategoryResponse, error) {
	// 查询一级分类
	query := `SELECT id, name, parent_id, sort_order, created_at FROM course_categories WHERE parent_id IS NULL ORDER BY sort_order ASC`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*models.CourseCategoryResponse
	for rows.Next() {
		var category models.CourseCategoryResponse
		var parentID sql.NullInt64
		
		if err := rows.Scan(&category.ID, &category.Name, &parentID, &category.SortOrder, &category.CreatedAt); err != nil {
			return nil, err
		}
		
		if parentID.Valid {
			category.ParentID = &parentID.Int64
		}
		
		// 查询子分类
		children, err := s.getChildrenCategories(category.ID)
		if err != nil {
			return nil, err
		}
		category.Children = children
		
		categories = append(categories, &category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

// getChildrenCategories 获取子分类
func (s *courseCategoryService) getChildrenCategories(parentID int64) ([]*models.CourseCategoryResponse, error) {
	query := `SELECT id, name, parent_id, sort_order, created_at FROM course_categories WHERE parent_id = ? ORDER BY sort_order ASC`
	rows, err := s.db.Query(query, parentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var children []*models.CourseCategoryResponse
	for rows.Next() {
		var child models.CourseCategoryResponse
		var childParentID sql.NullInt64
		
		if err := rows.Scan(&child.ID, &child.Name, &childParentID, &child.SortOrder, &child.CreatedAt); err != nil {
			return nil, err
		}
		
	if childParentID.Valid {
		child.ParentID = &childParentID.Int64
	}
	
	// 递归获取子分类
	grandChildren, err := s.getChildrenCategories(child.ID)
	if err != nil {
		return nil, err
	}
	child.Children = grandChildren
	
	children = append(children, &child)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return children, nil
}

// GetCategoryByID 根据ID获取课程分类
func (s *courseCategoryService) GetCategoryByID(id int64) (*models.CourseCategoryResponse, error) {
	query := `SELECT id, name, parent_id, sort_order, created_at FROM course_categories WHERE id = ?`
	row := s.db.QueryRow(query, id)

	var category models.CourseCategoryResponse
	var parentID sql.NullInt64
	
	if err := row.Scan(&category.ID, &category.Name, &parentID, &category.SortOrder, &category.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("分类不存在")
		}
		return nil, err
	}

	if parentID.Valid {
		category.ParentID = &parentID.Int64
	}

	// 获取子分类
	children, err := s.getChildrenCategories(category.ID)
	if err != nil {
		return nil, err
	}
	category.Children = children

	return &category, nil
}

// CreateCategory 创建课程分类
func (s *courseCategoryService) CreateCategory(category *models.CourseCategory) error {
	query := `INSERT INTO course_categories (name, parent_id, sort_order, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
	now := time.Now()

	result, err := s.db.Exec(query, category.Name, category.ParentID, category.SortOrder, now, now)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	category.ID = id

	return nil
}

// UpdateCategory 更新课程分类
func (s *courseCategoryService) UpdateCategory(category *models.CourseCategory) error {
	query := `UPDATE course_categories SET name = ?, parent_id = ?, sort_order = ?, updated_at = ? WHERE id = ?`
	now := time.Now()

	result, err := s.db.Exec(query, category.Name, category.ParentID, category.SortOrder, now, category.ID)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("分类不存在")
	}

	return nil
}

// DeleteCategory 删除课程分类
func (s *courseCategoryService) DeleteCategory(id int64) error {
	// 检查是否有子分类
	var count int
	query := `SELECT COUNT(*) FROM course_categories WHERE parent_id = ?`
	if err := s.db.QueryRow(query, id).Scan(&count); err != nil {
		return err
	}

	if count > 0 {
		return errors.New("该分类下有子分类，无法删除")
	}

	// 检查是否有课程使用该分类
	query = `SELECT COUNT(*) FROM courses WHERE category_id = ?`
	if err := s.db.QueryRow(query, id).Scan(&count); err != nil {
		return err
	}

	if count > 0 {
		return errors.New("该分类下有课程，无法删除")
	}

	// 删除分类
	query = `DELETE FROM course_categories WHERE id = ?`
	result, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("分类不存在")
	}

	return nil
}