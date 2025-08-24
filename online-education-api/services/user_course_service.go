package services

import (
	"database/sql"
	"errors"
	"online-education-api/models"
	"time"
)

// UserCourseService 用户课程服务接口
type UserCourseService interface {
	EnrollCourse(userID, courseID int64) error
	GetUserCourses(userID int64, page, pageSize int) ([]*models.UserCourseResponse, int, error)
	GetUserCourseByID(userID, courseID int64) (*models.UserCourseResponse, error)
	UnenrollCourse(userID, courseID int64) error
}

// userCourseService 用户课程服务实现
 type userCourseService struct {
	db *sql.DB
}

// NewUserCourseService 创建用户课程服务实例
func NewUserCourseService(db *sql.DB) UserCourseService {
	return &userCourseService{db: db}
}

// EnrollCourse 用户报名课程
func (s *userCourseService) EnrollCourse(userID, courseID int64) error {
	// 检查课程是否存在
	courseService := NewCourseService(s.db)
	course, err := courseService.GetCourseDetail(courseID)
	if err != nil {
		return err
	}

	// 检查用户是否已报名该课程
	var exists int
	query := `SELECT COUNT(*) FROM user_courses WHERE user_id = ? AND course_id = ? AND status = 1`
	if err := s.db.QueryRow(query, userID, courseID).Scan(&exists); err != nil {
		return err
	}

	if exists > 0 {
		return errors.New("您已报名该课程")
	}

	// 创建订单ID (简化版，实际应使用更复杂的生成规则)
	orderID := "ORD" + time.Now().Format("20060102150405") + string(rune(userID%10000)) + string(rune(courseID%10000))

	// 插入用户课程关系
	query = `INSERT INTO user_courses (user_id, course_id, order_id, price, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)`
	now := time.Now()

	_, err = s.db.Exec(query, userID, courseID, orderID, course.Price, 1, now, now)
	if err != nil {
		return err
	}

	// 更新课程学生数量
	query = `UPDATE courses SET student_count = student_count + 1 WHERE id = ?`
	_, err = s.db.Exec(query, courseID)
	if err != nil {
		return err
	}

	return nil
}

// GetUserCourses 获取用户报名的课程列表
func (s *userCourseService) GetUserCourses(userID int64, page, pageSize int) ([]*models.UserCourseResponse, int, error) {
	// 计算偏移量
	offset := (page - 1) * pageSize

	// 查询用户课程关系
	query := `SELECT uc.id, uc.course_id, uc.price, uc.status, uc.created_at, c.title, c.cover_image, c.teacher_id, c.level, c.duration FROM user_courses uc LEFT JOIN courses c ON uc.course_id = c.id WHERE uc.user_id = ? AND uc.status = 1 ORDER BY uc.created_at DESC LIMIT ? OFFSET ?`
	rows, err := s.db.Query(query, userID, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// 获取总记录数
	var total int
	countQuery := `SELECT COUNT(*) FROM user_courses WHERE user_id = ? AND status = 1`
	if err := s.db.QueryRow(countQuery, userID).Scan(&total); err != nil {
		return nil, 0, err
	}

	// 解析结果
	var userCourses []*models.UserCourseResponse
	for rows.Next() {
		var userCourse models.UserCourseResponse
		var course models.Course
		
		if err := rows.Scan(&userCourse.ID, &userCourse.CourseID, &userCourse.Price, &userCourse.Status, &userCourse.CreatedAt, &course.Title, &course.CoverImage, &course.TeacherID, &course.Level, &course.Duration); err != nil {
			return nil, 0, err
		}
		
		course.ID = userCourse.CourseID
		userCourse.Course = &course
		
		userCourses = append(userCourses, &userCourse)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return userCourses, total, nil
}

// GetUserCourseByID 检查用户是否报名了某课程
func (s *userCourseService) GetUserCourseByID(userID, courseID int64) (*models.UserCourseResponse, error) {
	query := `SELECT uc.id, uc.course_id, uc.price, uc.status, uc.created_at FROM user_courses uc WHERE uc.user_id = ? AND uc.course_id = ?`
	row := s.db.QueryRow(query, userID, courseID)

	var userCourse models.UserCourseResponse
	
	if err := row.Scan(&userCourse.ID, &userCourse.CourseID, &userCourse.Price, &userCourse.Status, &userCourse.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("未找到该课程报名记录")
		}
		return nil, err
	}

	return &userCourse, nil
}

// UnenrollCourse 用户取消报名课程
func (s *userCourseService) UnenrollCourse(userID, courseID int64) error {
	// 检查是否存在该报名记录
	query := `SELECT COUNT(*) FROM user_courses WHERE user_id = ? AND course_id = ? AND status = 1`
	var exists int
	if err := s.db.QueryRow(query, userID, courseID).Scan(&exists); err != nil {
		return err
	}

	if exists == 0 {
		return errors.New("未找到该课程报名记录")
	}

	// 更新状态为已退款
	query = `UPDATE user_courses SET status = 0, updated_at = ? WHERE user_id = ? AND course_id = ?`
	now := time.Now()
	result, err := s.db.Exec(query, now, userID, courseID)
	if err != nil {
		return err
	}

	// 更新课程学生数量
	query = `UPDATE courses SET student_count = student_count - 1 WHERE id = ?`
	_, err = s.db.Exec(query, courseID)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("取消报名失败")
	}

	return nil
}