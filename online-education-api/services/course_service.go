package services

import (
	"database/sql"
	"errors"
	"online-education-api/models"
	"time"
)

// CourseService 课程服务接口
type CourseService interface {
	GetCourseList(page, pageSize int, categoryID, level int64, search string) ([]*models.CourseResponse, int, error)
	GetCourseDetail(id int64) (*models.CourseDetailResponse, error)
	CreateCourse(course *models.Course) error
	UpdateCourse(course *models.Course) error
	DeleteCourse(id int64) error
	GetCoursesByCategory(categoryID int64, page, pageSize int) ([]*models.CourseResponse, int, error)
}

// courseService 课程服务实现
 type courseService struct {
	db *sql.DB
}

// NewCourseService 创建课程服务实例
func NewCourseService(db *sql.DB) CourseService {
	return &courseService{db: db}
}

// GetCourseList 获取课程列表
func (s *courseService) GetCourseList(page, pageSize int, categoryID, level int64, search string) ([]*models.CourseResponse, int, error) {
	// 计算偏移量
	offset := (page - 1) * pageSize

	// 构建查询条件
	query := `SELECT c.id, c.title, c.description, c.cover_image, c.price, c.original_price, c.category_id, c.teacher_id, c.level, c.duration, c.student_count, c.rating, c.status, c.created_at, cc.name as category_name, u.username as teacher_name FROM courses c LEFT JOIN course_categories cc ON c.category_id = cc.id LEFT JOIN users u ON c.teacher_id = u.id WHERE 1=1`
	countQuery := `SELECT COUNT(*) FROM courses c WHERE 1=1`

	// 添加查询参数
	params := []interface{}{}

	if categoryID > 0 {
		query += ` AND c.category_id = ?`
		countQuery += ` AND category_id = ?`
		params = append(params, categoryID)
	}

	if level > 0 {
		query += ` AND c.level = ?`
		countQuery += ` AND level = ?`
		params = append(params, level)
	}

	if search != "" {
		query += ` AND (c.title LIKE ? OR c.description LIKE ?)`
		countQuery += ` AND (title LIKE ? OR description LIKE ?)`
		searchParam := "%" + search + "%"
		params = append(params, searchParam, searchParam)
	}

	// 添加排序和分页
	query += ` ORDER BY c.created_at DESC LIMIT ? OFFSET ?`
	params = append(params, pageSize, offset)

	// 执行查询
	rows, err := s.db.Query(query, params...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// 获取总记录数
	var total int
	countParams := params[:len(params)-2] // 移除limit和offset参数
	if err := s.db.QueryRow(countQuery, countParams...).Scan(&total); err != nil {
		return nil, 0, err
	}

	// 解析结果
	var courses []*models.CourseResponse
	for rows.Next() {
		var course models.CourseResponse
		
		if err := rows.Scan(&course.ID, &course.Title, &course.Description, &course.CoverImage, &course.Price, &course.OriginalPrice, &course.CategoryID, &course.TeacherID, &course.Level, &course.Duration, &course.StudentCount, &course.Rating, &course.Status, &course.CreatedAt, &course.CategoryName, &course.TeacherName); err != nil {
			return nil, 0, err
		}
		
		courses = append(courses, &course)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return courses, total, nil
}

// GetCourseDetail 获取课程详情
func (s *courseService) GetCourseDetail(id int64) (*models.CourseDetailResponse, error) {
	// 查询课程基本信息
	query := `SELECT c.id, c.title, c.description, c.cover_image, c.price, c.original_price, c.category_id, c.teacher_id, c.level, c.duration, c.student_count, c.rating, c.status, c.created_at, cc.name as category_name, u.username as teacher_name, u.avatar as teacher_avatar FROM courses c LEFT JOIN course_categories cc ON c.category_id = cc.id LEFT JOIN users u ON c.teacher_id = u.id WHERE c.id = ?`
	row := s.db.QueryRow(query, id)

	var course models.Course
	var categoryName, teacherName, teacherAvatar string
	
	if err := row.Scan(&course.ID, &course.Title, &course.Description, &course.CoverImage, &course.Price, &course.OriginalPrice, &course.CategoryID, &course.TeacherID, &course.Level, &course.Duration, &course.StudentCount, &course.Rating, &course.Status, &course.CreatedAt, &categoryName, &teacherName, &teacherAvatar); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("课程不存在")
		}
		return nil, err
	}

	// 查询分类信息
	category, err := NewCourseCategoryService(s.db).GetCategoryByID(course.CategoryID)
	if err != nil {
		return nil, err
	}

	// 查询教师信息
	teacher := &models.UserResponse{
		ID:       course.TeacherID,
		Username: teacherName,
		Avatar:   teacherAvatar,
	}

	// 查询章节和课时信息
	chapters, totalLessons, err := s.getCourseChaptersAndLessons(id)
	if err != nil {
		return nil, err
	}

	// 构建响应
	courseDetail := &models.CourseDetailResponse{
		Course:      course,
		Category:    category,
		Teacher:     teacher,
		Chapters:    chapters,
		TotalLessons: totalLessons,
	}

	return courseDetail, nil
}

// getCourseChaptersAndLessons 获取课程的章节和课时
func (s *courseService) getCourseChaptersAndLessons(courseID int64) ([]*models.Chapter, int, error) {
	// 查询章节
	query := `SELECT id, course_id, title, sort_order, created_at FROM chapters WHERE course_id = ? ORDER BY sort_order ASC`
	rows, err := s.db.Query(query, courseID)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var chapters []*models.Chapter
	totalLessons := 0

	for rows.Next() {
		var chapter models.Chapter
		
		if err := rows.Scan(&chapter.ID, &chapter.CourseID, &chapter.Title, &chapter.SortOrder, &chapter.CreatedAt); err != nil {
			return nil, 0, err
		}

		// 查询课时
		lessonsQuery := `SELECT id, chapter_id, title, video_url, duration, sort_order, free, created_at FROM lessons WHERE chapter_id = ? ORDER BY sort_order ASC`
		lessonsRows, err := s.db.Query(lessonsQuery, chapter.ID)
		if err != nil {
			return nil, 0, err
		}
		defer lessonsRows.Close()

		var lessons []*models.Lesson
		for lessonsRows.Next() {
			var lesson models.Lesson
			
			if err := lessonsRows.Scan(&lesson.ID, &lesson.ChapterID, &lesson.Title, &lesson.VideoURL, &lesson.Duration, &lesson.SortOrder, &lesson.Free, &lesson.CreatedAt); err != nil {
				return nil, 0, err
			}
			
			lessons = append(lessons, &lesson)
		}

		if err = lessonsRows.Err(); err != nil {
			return nil, 0, err
		}

		chapter.Lessons = lessons
		totalLessons += len(lessons)
		chapters = append(chapters, &chapter)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return chapters, totalLessons, nil
}

// CreateCourse 创建课程
func (s *courseService) CreateCourse(course *models.Course) error {
	query := `INSERT INTO courses (title, description, cover_image, price, original_price, category_id, teacher_id, level, duration, student_count, rating, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	now := time.Now()

	// 默认值
	if course.StudentCount == 0 {
		course.StudentCount = 0
	}
	if course.Rating == 0 {
		course.Rating = 0
	}

	result, err := s.db.Exec(query, course.Title, course.Description, course.CoverImage, course.Price, course.OriginalPrice, course.CategoryID, course.TeacherID, course.Level, course.Duration, course.StudentCount, course.Rating, course.Status, now, now)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	course.ID = id

	return nil
}

// UpdateCourse 更新课程
func (s *courseService) UpdateCourse(course *models.Course) error {
	query := `UPDATE courses SET title = ?, description = ?, cover_image = ?, price = ?, original_price = ?, category_id = ?, level = ?, duration = ?, status = ?, updated_at = ? WHERE id = ?`
	now := time.Now()

	result, err := s.db.Exec(query, course.Title, course.Description, course.CoverImage, course.Price, course.OriginalPrice, course.CategoryID, course.Level, course.Duration, course.Status, now, course.ID)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("课程不存在")
	}

	return nil
}

// DeleteCourse 删除课程
func (s *courseService) DeleteCourse(id int64) error {
	// 检查是否有章节关联
	var count int
	query := `SELECT COUNT(*) FROM chapters WHERE course_id = ?`
	if err := s.db.QueryRow(query, id).Scan(&count); err != nil {
		return err
	}

	if count > 0 {
		return errors.New("该课程下有关联章节，无法删除")
	}

	// 删除课程
	query = `DELETE FROM courses WHERE id = ?`
	result, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("课程不存在")
	}

	return nil
}

// GetCoursesByCategory 根据分类ID获取课程列表
func (s *courseService) GetCoursesByCategory(categoryID int64, page, pageSize int) ([]*models.CourseResponse, int, error) {
	return s.GetCourseList(page, pageSize, categoryID, 0, "")
}