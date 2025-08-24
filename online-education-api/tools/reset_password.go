package tools

import (
	"fmt"
	"log"

	"online-education-api/config"
	"golang.org/x/crypto/bcrypt"
)

// ResetPassword 重置管理员密码
func ResetPassword() {
	// 获取数据库配置
	dbConfig := config.GetDBConfig()

	// 初始化数据库连接
	db, err := config.InitDB(dbConfig)
	if err != nil {
		log.Fatalf("无法初始化数据库: %v", err)
	}
	defer db.Close()

	// 新的管理员密码
	newPassword := "admin123"

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("加密密码失败: %v", err)
	}

	// 更新管理员密码
	query := "UPDATE users SET password = ? WHERE username = 'admin'"
	result, err := db.Exec(query, string(hashedPassword))
	if err != nil {
		log.Fatalf("更新密码失败: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("获取受影响行数失败: %v", err)
	}

	fmt.Printf("成功更新 %d 行数据。新的管理员密码为: %s\n", rowsAffected, newPassword)
}

// 可以通过命令行调用此工具
// func main() {
//	ResetPassword()
// }