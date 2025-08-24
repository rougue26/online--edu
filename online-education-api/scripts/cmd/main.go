package main

import (
	"fmt"
	"log"

	"online-education-api/scripts/admin"
)

func main() {
	err := admin.CreateAdmin()
	if err != nil {
		log.Fatalf("创建管理员账户失败: %v", err)
	}

	fmt.Println("管理员账户创建流程已完成!")
}