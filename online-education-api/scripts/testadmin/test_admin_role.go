package testadmin

import (
	"fmt"
	"log"

	"online-education-api/utils"
)

// TestToken 测试JWT令牌中的角色信息
func TestToken(token string) {
	if token == "" {
		log.Println("请提供有效的JWT令牌")
		return
	}

	// 解析令牌
	claims, err := utils.ParseToken(token)
	if err != nil {
		log.Fatalf("解析令牌失败: %v", err)
	}

	// 打印角色信息
	fmt.Printf("用户ID: %d\n", claims.UserID)
	fmt.Printf("用户名: %s\n", claims.Username)
	fmt.Printf("角色: %s\n", claims.Role)

	// 检查是否为管理员
	if claims.Role == "admin" {
		fmt.Println("验证成功: 这是一个管理员账户")
	} else {
		fmt.Println("验证失败: 这不是一个管理员账户")
	}
}