package admin

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"online-education-api/config"
)

// 生成随机密码
func generateRandomPassword(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	// 生成字母数字密码
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}

// 创建管理员账户
func createAdmin(username, email, password string) error {
	// 获取数据库配置
	dbConfig := config.GetDBConfig()

	// 初始化数据库连接
	db, err := config.InitDB(dbConfig)
	if err != nil {
		return fmt.Errorf("无法初始化数据库: %v", err)
	}
	defer db.Close()

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("加密密码失败: %v", err)
	}

	// 插入管理员账户
	stmt := `INSERT INTO users (username, email, password, role)
		SELECT ?, ?, ?, 'admin'
		WHERE NOT EXISTS (SELECT 1 FROM users WHERE username = ? OR email = ?)`

	_, err = db.Exec(stmt, username, email, string(hashedPassword), username, email)
	if err != nil {
		return fmt.Errorf("执行SQL语句失败: %v", err)
	}

	return nil
}

func CreateAdmin() error {
	// 默认管理员信息
	username := "admin_new"
	email := "admin_new@example.com"
	password := "Password123!"

	// 或者生成随机密码
	// randomPassword, err := generateRandomPassword(12)
	// if err != nil {
	//	return fmt.Errorf("生成随机密码失败: %v", err)
	// }
	// password = randomPassword

	// 创建管理员账户
	err := createAdmin(username, email, password)
	if err != nil {
		return fmt.Errorf("创建管理员账户失败: %v", err)
	}

	fmt.Printf("管理员账户创建成功!\n用户名: %s\n邮箱: %s\n密码: %s\n", username, email, password)
	return nil
}