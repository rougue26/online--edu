package utils

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTSecret JWT密钥，实际应用中应从环境变量或配置文件中读取
var JWTSecret = []byte("1234567890abcdef1234567890abcdef") // 更新为更安全的密钥

// Claims 自定义JWT声明
 type Claims struct {
	UserID int64  `json:"user_id"`
	Username string `json:"username"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT令牌
func GenerateToken(userID int64, username string, role string) (string, error) {
	// 设置令牌过期时间
	expirationTime := time.Now().Add(24 * time.Hour) // 24小时后过期

	// 创建声明
	claims := &Claims{
		UserID: userID,
		Username: username,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	// 创建令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名令牌
	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		log.Printf("生成令牌失败: %v\n", err)
		return "", err
	}

	log.Printf("成功生成令牌 for user: %s\n", username)
	return tokenString, nil
}

// ParseToken 解析JWT令牌
func ParseToken(tokenString string) (*Claims, error) {
	// 解析令牌
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Printf("无效的签名算法: %v\n", token.Header["alg"])
			return nil, errors.New("无效的签名算法")
		}
		return JWTSecret, nil
	})

	if err != nil {
		log.Printf("解析令牌失败: %v\n", err)
		return nil, err
	}

	// 验证令牌是否有效
	if !token.Valid {
		log.Println("无效的令牌")
		return nil, errors.New("无效的令牌")
	}

	log.Printf("成功解析令牌 for user: %s\n", claims.Username)
	return claims, nil
}