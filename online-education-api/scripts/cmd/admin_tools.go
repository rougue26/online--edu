package main

import (
	"fmt"
	"os"

	"online-education-api/scripts/checkadmin"
)

func main() {
	// 检查命令行参数
	if len(os.Args) < 2 {
		fmt.Println("用法: admin_tools <command>")
		fmt.Println("可用命令:")
		fmt.Println("  check-role    检查并更新管理员角色")
		os.Exit(1)
	}

	// 根据命令调用相应的函数
	switch os.Args[1] {
	case "check-role":
		checkadmin.CheckAndUpdateAdminRole()
	default:
		fmt.Println("未知命令")
		os.Exit(1)
	}
}