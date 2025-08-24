package main

import (
	"fmt"
	"log"

	"online-education-api/utils"
)

func main() {
	// 测试JWT令牌生成和验证
	// 假设我们有一个管理员用户
	userID := int64(1)
	username := "admin"
	role := "admin"

	// 生成JWT令牌
	token, err := utils.GenerateToken(userID, username, role)
	if err != nil {
		log.Fatalf("生成令牌失败: %v", err)
	}

	fmt.Printf("生成的JWT令牌: %s\n\n", token)

	// 解析JWT令牌
	claims, err := utils.ParseToken(token)
	if err != nil {
		log.Fatalf("解析令牌失败: %v", err)
	}

	// 打印令牌中的信息
	fmt.Println("令牌中的信息:")
	fmt.Printf("用户ID: %d\n", claims.UserID)
	fmt.Printf("用户名: %s\n", claims.Username)
	fmt.Printf("角色: %s\n\n", claims.Role)

	// 验证角色
	if claims.Role == "admin" {
		fmt.Println("验证成功: 令牌中包含管理员角色")
	} else {
		fmt.Println("验证失败: 令牌中不包含管理员角色")
	}
}