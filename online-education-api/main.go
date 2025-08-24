package main

import (
	"log"
	"net/http"

	"online-education-api/config"
	"online-education-api/controllers"
	"online-education-api/middleware"
	"online-education-api/routes"
	"online-education-api/services"
)

func main() {
	// 获取数据库配置
	dbConfig := config.GetDBConfig()

	// 初始化数据库连接
	db, err := config.InitDB(dbConfig)
	if err != nil {
		log.Fatalf("无法初始化数据库: %v", err)
	}
	defer db.Close()

	// 创建服务实例
	videoService := services.NewVideoService(db)
	userService := services.NewUserService(db)
	courseCategoryService := services.NewCourseCategoryService(db)
	courseService := services.NewCourseService(db)
	userCourseService := services.NewUserCourseService(db)
	postService := services.NewPostService(db)
	paymentService := services.NewPaymentService(db)
	commentService := services.NewCommentService(db)

	// 创建控制器实例
	videoController := controllers.NewVideoController(videoService)
	userController := controllers.NewUserController(userService)
	courseCategoryController := controllers.NewCourseCategoryController(courseCategoryService)
	courseController := controllers.NewCourseController(courseService)
	userCourseController := controllers.NewUserCourseController(userCourseService)
	postController := controllers.NewPostController(postService)
	paymentController := controllers.NewPaymentController(paymentService)
	commentController := controllers.NewCommentController(commentService)

	// 设置路由
	r := routes.SetupRoutes(videoController, userController, courseCategoryController, courseController, userCourseController, postController, paymentController, commentController)

	// 应用CORS中间件
	log.Println("服务器启动在 http://localhost:8082")
	if err := http.ListenAndServe(":8082", middleware.CORSMiddleware(r)); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}