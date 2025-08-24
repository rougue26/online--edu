package checkadmin

import (
	"fmt"
	"log"

	"online-education-api/config"
)

// CheckAndUpdateAdminRole 检查并更新管理员角色
func CheckAndUpdateAdminRole() {
	// 获取数据库配置
	dbConfig := config.GetDBConfig()

	// 初始化数据库连接
	db, err := config.InitDB(dbConfig)
	if err != nil {
		log.Fatalf("无法初始化数据库: %v", err)
	}
	defer db.Close()

	// 查询管理员账户
	var username, email, role string
	query := "SELECT username, email, role FROM users WHERE username = ?"
	err = db.QueryRow(query, "admin").Scan(&username, &email, &role)
	if err != nil {
		log.Fatalf("查询管理员账户失败: %v", err)
	}

	fmt.Printf("管理员账户信息:\n")
	fmt.Printf("用户名: %s\n", username)
	fmt.Printf("邮箱: %s\n", email)
	fmt.Printf("角色: %s\n", role)

	// 确保角色是admin
	if role != "admin" {
		fmt.Printf("警告: 管理员账户角色不是'admin'，正在更新...\n")
		updateQuery := "UPDATE users SET role = 'admin' WHERE username = ?"
		_, err = db.Exec(updateQuery, "admin")
		if err != nil {
			log.Fatalf("更新管理员角色失败: %v", err)
		}
		fmt.Printf("管理员角色已更新为'admin'\n")
	} else {
		fmt.Printf("管理员角色验证成功\n")
	}
}