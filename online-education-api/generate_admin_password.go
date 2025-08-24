package main

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

func main() {
	// 设置新的admin密码
	password := "admin123"

	// 生成密码哈希
	hash, err := GeneratePasswordHash(password)
	if err != nil {
		panic(err)
	}

	// 打印密码哈希
	fmt.Println("密码:", password)
	fmt.Println("密码哈希:", hash)

	// 输出更新数据库的SQL语句
	fmt.Println(`
更新admin用户密码的SQL语句:`)
	fmt.Printf("UPDATE users SET password = '%s' WHERE username = 'admin';\n", hash)
}