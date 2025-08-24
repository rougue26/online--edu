package middleware

import (
	"context"
	"net/http"
	"strings"

	"online-education-api/utils"
)

// AuthMiddleware JWT认证中间件
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 从请求头中获取Authorization头
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "未提供认证令牌", http.StatusUnauthorized)
			return
		}

		// 检查令牌格式
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "无效的认证令牌格式", http.StatusUnauthorized)
			return
		}

		// 解析令牌
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			http.Error(w, "无效的认证令牌: " + err.Error(), http.StatusUnauthorized)
			return
		}

		// 将用户信息和角色添加到请求上下文中
	ctx := context.WithValue(r.Context(), "userID", claims.UserID)
	ctx = context.WithValue(ctx, "username", claims.Username)
	ctx = context.WithValue(ctx, "role", claims.Role)

		// 继续处理请求
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}