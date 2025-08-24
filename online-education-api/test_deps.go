package tools

import (
	"fmt"
	"log"

	"online-education-api/config"
	"online-education-api/models"
	"online-education-api/services"
	"online-education-api/utils"
)

// TestDeps 测试依赖项
func TestDeps() {
	// 测试数据库连接
	fmt.Println("Testing database connection...")
	dbConfig := config.GetDBConfig()
	var err error
	if err != nil {
		log.Fatalf("Failed to load database config: %v", err)
	}

	db, err := config.InitDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()
	fmt.Println("Database connection successful!")

	// 测试JWT工具
	fmt.Println("\nTesting JWT utility...")
	userID := int64(1)
	username := "testuser"
	token, err := utils.GenerateToken(userID, username)
	if err != nil {
		log.Fatalf("Failed to generate JWT token: %v", err)
	}
	fmt.Printf("Generated JWT token: %s\n", token)

	claims, err := utils.ParseToken(token)
	if err != nil {
		log.Fatalf("Failed to parse JWT token: %v", err)
	}
	fmt.Printf("Parsed token claims: UserID=%d, Username=%s\n", claims.UserID, claims.Username)

	// 测试服务初始化
	fmt.Println("\nTesting service initialization...")
	_ = services.NewUserService(db)
	_ = services.NewCourseService(db)
	_ = services.NewPostService(db)
	fmt.Println("All services initialized successfully!")

	// 打印测试完成信息
	fmt.Println("\nAll dependencies test completed successfully!")
	fmt.Println("You can run this test with: go run test_deps.go")

// 为了避免与main.go中的main函数冲突，这里不使用init函数
// 可以通过命令行调用此工具
// func main() {
//	TestDeps()
// }