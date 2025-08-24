package tools

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// GeneratePasswordHash 生成密码哈希
func GeneratePasswordHash(password string) (string, error) {
	// 生成密码哈希
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// 为了保持向后兼容性，保留main函数但添加注释说明这只是用于测试
// func main() {
//	password := "password"
//	hash, err := GeneratePasswordHash(password)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("密码哈希:", hash)
// }