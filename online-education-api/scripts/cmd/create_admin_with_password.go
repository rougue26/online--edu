package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	"online-education-api/config"
)

// CreateAdminWithPassword 创建或更新具有指定密码的管理员账户
func CreateAdminWithPassword(username, email, password string) error {
	role := "admin"

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

	// 插入或更新管理员账户
	// 首先尝试插入
	insertStmt := `INSERT INTO users (username, email, password, role)
		SELECT ?, ?, ?, ?
		WHERE NOT EXISTS (SELECT 1 FROM users WHERE username = ?)`

	result, err := db.Exec(insertStmt, username, email, string(hashedPassword), role, username)
	if err != nil {
		return fmt.Errorf("执行插入SQL语句失败: %v", err)
	}

	// 检查是否插入了新行
	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("获取影响行数失败: %v", err)
	}

	if affected == 0 {
		// 如果没有插入新行，说明用户已存在，更新密码
		updateStmt := `UPDATE users
		SET password = ?, email = ?, role = ?
		WHERE username = ?`

		_, err = db.Exec(updateStmt, string(hashedPassword), email, role, username)
		if err != nil {
			return fmt.Errorf("执行更新SQL语句失败: %v", err)
		}

		fmt.Printf("管理员账户已存在，密码已更新为: %s\n", password)
	} else {
		fmt.Printf("管理员账户创建成功!\n用户名: %s\n密码: %s\n", username, password)
	}

	return nil
}

func main() {
	// 默认管理员信息
	username := "admin"
	email := "admin@example.com"
	password := "abc123"

	err := CreateAdminWithPassword(username, email, password)
	if err != nil {
		log.Fatalf("创建/更新管理员账户失败: %v", err)
	}
}