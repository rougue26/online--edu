package main

import (
	"fmt"
	"log"

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

	// 检查cover_image_url列是否存在
	var columnExists bool
	query := "SELECT COUNT(*) FROM information_schema.columns WHERE table_schema = DATABASE() AND table_name = 'videos' AND column_name = 'cover_image_url'"
	err = db.QueryRow(query).Scan(&columnExists)
	if err != nil {
		log.Fatalf("检查列是否存在失败: %v", err)
	}

	// 如果列不存在，则添加
	if !columnExists {
		addColumnQuery := "ALTER TABLE videos ADD COLUMN cover_image_url VARCHAR(255)"
		_, err = db.Exec(addColumnQuery)
		if err != nil {
			log.Fatalf("添加cover_image_url列失败: %v", err)
		}
		fmt.Println("成功添加cover_image_url列!")
	} else {
		fmt.Println("cover_image_url列已存在，无需添加!")
	}
}