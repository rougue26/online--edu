package services

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"online-education-api/models"
	"online-education-api/utils"
	"golang.org/x/crypto/bcrypt"
)

// UserService 用户服务接口
type UserService interface {
	Register(user *models.UserRegisterRequest) (*models.User, error)
	Login(loginReq *models.UserLoginRequest) (*models.User, string, error)
	GetUserByID(id int64) (*models.User, error)
	UpdateUser(id int64, updateReq *models.UserUpdateRequest) (*models.User, error)
	ChangePassword(userID int64, oldPassword, newPassword string) error
	GetUserList(page, pageSize int) ([]*models.User, int64, error)
	CreateUser(user *models.UserCreateRequest) (*models.User, error)
	DeleteUser(id int64) error
}

// userService 实现UserService接口
type userService struct {
	db *sql.DB
}

// NewUserService 创建用户服务实例
func NewUserService(db *sql.DB) UserService {
	return &userService{db: db}
}

// Register 注册新用户
func (s *userService) Register(user *models.UserRegisterRequest) (*models.User, error) {
	// 检查用户名是否已存在
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)"
	err := s.db.QueryRow(query, user.Username).Scan(&exists)
	if err != nil {
		return nil, fmt.Errorf("查询用户名失败: %w", err)
	}
	if exists {
		return nil, errors.New("用户名已存在")
	}

	// 检查邮箱是否已存在
	query = "SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)"
	err = s.db.QueryRow(query, user.Email).Scan(&exists)
	if err != nil {
		return nil, fmt.Errorf("查询邮箱失败: %w", err)
	}
	if exists {
		return nil, errors.New("邮箱已被注册")
	}

	// 加密密码
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("加密密码失败: %w", err)
	}

	// 创建用户
	now := time.Now()
	query = "INSERT INTO users (username, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	result, err := s.db.Exec(query, user.Username, user.Email, string(passwordHash), now, now)
	if err != nil {
		return nil, fmt.Errorf("创建用户失败: %w", err)
	}

	// 获取新创建用户的ID
	userID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("获取用户ID失败: %w", err)
	}

	// 返回新创建的用户
	return &models.User{
		ID:        userID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// Login 用户登录
func (s *userService) Login(loginReq *models.UserLoginRequest) (*models.User, string, error) {
	// 实现登录逻辑
	// 1. 根据用户名查询用户
	var user models.User
	var passwordHash string
	query := "SELECT id, username, email, password, COALESCE(avatar, ''), COALESCE(bio, ''), created_at, updated_at, role FROM users WHERE username = ?"
	fmt.Printf("Login attempt for username: %s\n", loginReq.Username)
	err := s.db.QueryRow(query, loginReq.Username).Scan(
		&user.ID, &user.Username, &user.Email, &passwordHash, &user.Avatar, &user.Bio,
		&user.CreatedAt, &user.UpdatedAt, &user.Role,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("User not found: %s\n", loginReq.Username)
			return nil, "", errors.New("用户名或密码错误")
		}
		fmt.Printf("Database query error: %v\n", err)
		return nil, "", fmt.Errorf("查询用户失败: %w", err)
	}

	// 2. 暂时移除用户状态检查，因为数据库中没有status字段
	// if status == 0 {
	// 	return nil, "", errors.New("账户已被禁用")
	// }

	// 3. 验证密码
	fmt.Printf("User found: %s, ID: %d\n", user.Username, user.ID)
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(loginReq.Password))
	if err != nil {
		fmt.Printf("Password verification failed for user: %s\n", user.Username)
		return nil, "", errors.New("用户名或密码错误")
	}

	// 4. 生成JWT令牌，包含角色信息
	fmt.Printf("Password verified, generating token for user: %s\n", user.Username)
	token, err := utils.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		fmt.Printf("Token generation failed: %v\n", err)
		return nil, "", fmt.Errorf("生成令牌失败: %w", err)
	}
	fmt.Printf("Token generated successfully for user: %s\n", user.Username)

	// 5. 更新最后登录时间 (暂时注释，因为数据库中没有last_login字段)
	// updateQuery := "UPDATE users SET last_login = NOW() WHERE id = ?"
	// _, err = s.db.Exec(updateQuery, user.ID)
	// if err != nil {
	// 	// 记录警告但不阻止登录
	// 	fmt.Printf("更新最后登录时间失败: %v\n", err)
	// }

	return &user, token, nil
}

// GetUserByID 根据ID获取用户信息
func (s *userService) GetUserByID(id int64) (*models.User, error) {
	var user models.User
	query := "SELECT id, username, email, COALESCE(avatar, ''), COALESCE(bio, ''), created_at, updated_at FROM users WHERE id = ?"
	err := s.db.QueryRow(query, id).Scan(
		&user.ID, &user.Username, &user.Email, &user.Avatar, &user.Bio,
		&user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("用户不存在")
		}
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	return &user, nil
}

// UpdateUser 更新用户信息
func (s *userService) UpdateUser(id int64, updateReq *models.UserUpdateRequest) (*models.User, error) {
	// 实现更新用户信息逻辑
	query := "UPDATE users SET updated_at = ?"
	args := []interface{}{time.Now()}

	if updateReq.Nickname != "" {
		query += ", nickname = ?"
		args = append(args, updateReq.Nickname)
	}

	if updateReq.Avatar != "" {
		query += ", avatar = ?"
		args = append(args, updateReq.Avatar)
	}

	if updateReq.Bio != "" {
		query += ", bio = ?"
		args = append(args, updateReq.Bio)
	}

	query += " WHERE id = ?"
	args = append(args, id)

	_, err := s.db.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("更新用户信息失败: %w", err)
	}

	// 返回更新后的用户信息
	return s.GetUserByID(id)
}

// ChangePassword 修改密码
func (s *userService) ChangePassword(userID int64, oldPassword, newPassword string) error {
	// 实现修改密码逻辑
	var passwordHash string
	query := "SELECT password FROM users WHERE id = ?"
	err := s.db.QueryRow(query, userID).Scan(&passwordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("用户不存在")
		}
		return fmt.Errorf("查询密码失败: %w", err)
	}

	// 验证旧密码
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(oldPassword))
	if err != nil {
		return errors.New("旧密码错误")
	}

	// 加密新密码
	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("加密新密码失败: %w", err)
	}

	// 更新密码
	query = "UPDATE users SET password = ?, updated_at = ? WHERE id = ?"
	_, err = s.db.Exec(query, string(newPasswordHash), time.Now(), userID)
	if err != nil {
		return fmt.Errorf("更新密码失败: %w", err)
	}

	return nil
}

// GetUserList 获取用户列表
func (s *userService) GetUserList(page, pageSize int) ([]*models.User, int64, error) {
	// 计算偏移量
	offset := (page - 1) * pageSize
	
	// 获取用户列表
	query := "SELECT id, username, email, COALESCE(avatar, ''), COALESCE(bio, ''), created_at, updated_at, role FROM users LIMIT ? OFFSET ?"
	rows, err := s.db.Query(query, pageSize, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("查询用户列表失败: %w", err)
	}
	defer rows.Close()
	
	var users []*models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID, &user.Username, &user.Email, &user.Avatar,
			&user.Bio, &user.CreatedAt, &user.UpdatedAt, &user.Role,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("扫描用户数据失败: %w", err)
		}
		users = append(users, &user)
	}
	
	// 检查遍历过程中是否有错误
	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("遍历用户数据失败: %w", err)
	}
	
	// 获取用户总数
	var total int64
	countQuery := "SELECT COUNT(*) FROM users"
	err = s.db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("查询用户总数失败: %w", err)
	}
	
	return users, total, nil
}

// CreateUser 创建新用户
func (s *userService) CreateUser(user *models.UserCreateRequest) (*models.User, error) {
	// 检查用户名是否已存在
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)"
	err := s.db.QueryRow(query, user.Username).Scan(&exists)
	if err != nil {
		return nil, fmt.Errorf("查询用户名失败: %w", err)
	}
	if exists {
		return nil, errors.New("用户名已存在")
	}

	// 检查邮箱是否已存在
	query = "SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)"
	err = s.db.QueryRow(query, user.Email).Scan(&exists)
	if err != nil {
		return nil, fmt.Errorf("查询邮箱失败: %w", err)
	}
	if exists {
		return nil, errors.New("邮箱已被注册")
	}

	// 加密密码
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("加密密码失败: %w", err)
	}

	// 设置默认角色
	role := "user"
	if user.Role != "" {
		role = user.Role
	}

	// 创建用户
	now := time.Now()
	query = "INSERT INTO users (username, email, password, role, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := s.db.Exec(query, user.Username, user.Email, string(passwordHash), role, now, now)
	if err != nil {
		return nil, fmt.Errorf("创建用户失败: %w", err)
	}

	// 获取新创建用户的ID
	userID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("获取用户ID失败: %w", err)
	}

	// 返回新创建的用户
	return &models.User{
		ID:        userID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      role,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// DeleteUser 删除用户
func (s *userService) DeleteUser(id int64) error {
	// 检查用户是否存在
	_, err := s.GetUserByID(id)
	if err != nil {
		return err
	}

	// 删除用户
	query := "DELETE FROM users WHERE id = ?"
	_, err = s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("删除用户失败: %w", err)
	}

	return nil
}