package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"online-education-api/config"
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

	// 读取初始化脚本
	script, err := ioutil.ReadFile("scripts/init_db.sql")
	if err != nil {
		log.Fatalf("无法读取初始化脚本: %v", err)
	}

	// 分割SQL语句
	statements := strings.Split(string(script), ";")

	// 逐个执行SQL语句
	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" || strings.HasPrefix(stmt, "--") {
			continue
		}

		_, err = db.Exec(stmt)
		if err != nil {
			log.Printf("执行SQL语句失败: %v\n语句: %s", err, stmt)
			// 继续执行其他语句
		}
	}

	fmt.Println("数据库初始化完成!")
}