package main

import (
	"database/sql"
	"fmt"
	"log"

	"online-education-api/config"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 获取数据库配置
	dbConfig := config.GetDBConfig()
	// 确认数据库密码是否正确
	// dbConfig.Password = "your_actual_password"

	// 初始化数据库连接
	db, err := config.InitDB(dbConfig)
	if err != nil {
		log.Fatalf("无法初始化数据库: %v", err)
	}
	defer db.Close()

	// 查询所有用户
	var users []struct {
		ID       int64  `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Role     string `json:"role"`
	}
	rows, err := db.Query("SELECT id, username, email, role FROM users")
	if err != nil {
		log.Fatalf("查询用户失败: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
			Email    string `json:"email"`
			Role     string `json:"role"`
		}
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Role); err != nil {
			log.Fatalf("扫描用户数据失败: %v", err)
		}
		users = append(users, user)
	}

	// 打印用户信息
	fmt.Println("数据库中的用户信息:")
	for _, user := range users {
		fmt.Printf("ID: %d, 用户名: %s, 邮箱: %s, 角色: %s\n",
			user.ID, user.Username, user.Email, user.Role)
	}

	// 查询管理员用户密码哈希
	var admin struct {
		ID           int64  `json:"id"`
		Username     string `json:"username"`
		PasswordHash string `json:"password_hash"`
	}
	err = db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", "admin").
		Scan(&admin.ID, &admin.Username, &admin.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("未找到管理员用户")
		} else {
			log.Printf("查询管理员用户失败: %v\n", err)
		}
	} else {
		fmt.Printf("\n管理员用户信息:\nID: %d, 用户名: %s\n",
			admin.ID, admin.Username)
		fmt.Printf("密码哈希: %s\n", admin.PasswordHash)
	}
}