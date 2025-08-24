package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"online-education-api/controllers"
	"online-education-api/middleware"
)

// SetupRoutes 设置路由
func SetupRoutes(
	videoController *controllers.VideoController,
	userController *controllers.UserController,
	courseCategoryController *controllers.CourseCategoryController,
	courseController *controllers.CourseController,
	userCourseController *controllers.UserCourseController,
	postController *controllers.PostController,
	paymentController *controllers.PaymentController,
	commentController *controllers.CommentController,
) *mux.Router {
	// 创建路由器
	r := mux.NewRouter()

	// 视频路由
	videoRoutes := r.PathPrefix("/api/videos").Subrouter()
	videoRoutes.HandleFunc("", videoController.GetVideoList).Methods("GET")
	videoRoutes.HandleFunc("", videoController.CreateVideo).Methods("POST")
	videoRoutes.HandleFunc("/{id}", videoController.GetVideoByID).Methods("GET")
	videoRoutes.HandleFunc("/{id}", videoController.UpdateVideo).Methods("PUT")
	videoRoutes.HandleFunc("/{id}", videoController.DeleteVideo).Methods("DELETE")

	// 视频分类路由
	videoRoutes.HandleFunc("/categories", videoController.GetVideoCategories).Methods("GET")

	// 课程分类路由
	courseCategoryRoutes := r.PathPrefix("/api/course-categories").Subrouter()
	courseCategoryRoutes.HandleFunc("", courseCategoryController.GetAllCategories).Methods("GET")
	courseCategoryRoutes.HandleFunc("/{id}", courseCategoryController.GetCategoryByID).Methods("GET")

	// 受保护的课程分类路由
	protectedCourseCategoryRoutes := courseCategoryRoutes.PathPrefix("").Subrouter()
	protectedCourseCategoryRoutes.Use(middleware.AuthMiddleware)
	protectedCourseCategoryRoutes.HandleFunc("", courseCategoryController.CreateCategory).Methods("POST")
	protectedCourseCategoryRoutes.HandleFunc("/{id}", courseCategoryController.UpdateCategory).Methods("PUT")
	protectedCourseCategoryRoutes.HandleFunc("/{id}", courseCategoryController.DeleteCategory).Methods("DELETE")

	// 课程路由
	courseRoutes := r.PathPrefix("/api/courses").Subrouter()
	courseRoutes.HandleFunc("", courseController.GetCourseList).Methods("GET")
	courseRoutes.HandleFunc("/{id}", courseController.GetCourseDetail).Methods("GET")

	// 受保护的课程路由
	protectedCourseRoutes := courseRoutes.PathPrefix("").Subrouter()
	protectedCourseRoutes.Use(middleware.AuthMiddleware)
	protectedCourseRoutes.HandleFunc("", courseController.CreateCourse).Methods("POST")
	protectedCourseRoutes.HandleFunc("/{id}", courseController.UpdateCourse).Methods("PUT")
	protectedCourseRoutes.HandleFunc("/{id}", courseController.DeleteCourse).Methods("DELETE")

	// 用户课程路由
	userCourseRoutes := r.PathPrefix("/api/user-courses").Subrouter()
	userCourseRoutes.Use(middleware.AuthMiddleware)
	userCourseRoutes.HandleFunc("", userCourseController.GetUserCourses).Methods("GET")
	userCourseRoutes.HandleFunc("/{courseID}", userCourseController.EnrollCourse).Methods("POST")
	userCourseRoutes.HandleFunc("/{courseID}", userCourseController.GetUserCourseByID).Methods("GET")
	userCourseRoutes.HandleFunc("/{courseID}", userCourseController.UnenrollCourse).Methods("DELETE")

	// 社区帖子路由
	postRoutes := r.PathPrefix("/api/posts").Subrouter()
	postRoutes.HandleFunc("", postController.GetPostList).Methods("GET")
	postRoutes.HandleFunc("/{id}", postController.GetPostDetail).Methods("GET")

	// 受保护的帖子路由
	protectedPostRoutes := postRoutes.PathPrefix("").Subrouter()
	protectedPostRoutes.Use(middleware.AuthMiddleware)
	protectedPostRoutes.HandleFunc("", postController.CreatePost).Methods("POST")
	protectedPostRoutes.HandleFunc("/{id}", postController.UpdatePost).Methods("PUT")
	protectedPostRoutes.HandleFunc("/{id}", postController.DeletePost).Methods("DELETE")
	protectedPostRoutes.HandleFunc("/user/posts", postController.GetUserPosts).Methods("GET")

	// 用户相关路由
	userRoutes := r.PathPrefix("/api/users").Subrouter()
	userRoutes.HandleFunc("/register", userController.Register).Methods("POST")
	userRoutes.HandleFunc("/login", userController.Login).Methods("POST")

	// 受保护的用户路由
	protectedUserRoutes := userRoutes.PathPrefix("").Subrouter()
	protectedUserRoutes.Use(middleware.AuthMiddleware)
	protectedUserRoutes.HandleFunc("/profile", userController.GetProfile).Methods("GET")
	protectedUserRoutes.HandleFunc("/profile", userController.UpdateProfile).Methods("PUT")
	protectedUserRoutes.HandleFunc("/password", userController.ChangePassword).Methods("PUT")
	protectedUserRoutes.HandleFunc("", userController.GetUserList).Methods("GET")
	protectedUserRoutes.HandleFunc("", userController.CreateUser).Methods("POST")
	protectedUserRoutes.HandleFunc("/{id}", userController.GetUserByID).Methods("GET")
	protectedUserRoutes.HandleFunc("/{id}", userController.UpdateUser).Methods("PUT")
	protectedUserRoutes.HandleFunc("/{id}", userController.DeleteUser).Methods("DELETE")

	// 支付路由
paymentRoutes := r.PathPrefix("/api/payments").Subrouter()
paymentRoutes.HandleFunc("/status/{orderID}", paymentController.GetPaymentStatus).Methods("GET")
	paymentRoutes.HandleFunc("/notify", paymentController.NotifyPayment).Methods("POST")

	// 受保护的支付路由
	protectedPaymentRoutes := paymentRoutes.PathPrefix("").Subrouter()
	protectedPaymentRoutes.Use(middleware.AuthMiddleware)
	protectedPaymentRoutes.HandleFunc("", paymentController.CreatePayment).Methods("POST")
	protectedPaymentRoutes.HandleFunc("/user", paymentController.GetUserPayments).Methods("GET")

	// 评论路由
	var commentRoutes = r.PathPrefix("/api/comments").Subrouter()
	commentRoutes.HandleFunc("", commentController.GetCommentList).Methods("GET")

	// 受保护的评论路由
	var protectedCommentRoutes = commentRoutes.PathPrefix("").Subrouter()
	protectedCommentRoutes.Use(middleware.AuthMiddleware)
	protectedCommentRoutes.HandleFunc("/{id}", commentController.DeleteComment).Methods("DELETE")

	// 首页路由
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Online Education API is running!"))
	}).Methods("GET")

	return r
}