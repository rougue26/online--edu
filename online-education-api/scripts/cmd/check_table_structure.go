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

	// 检查videos表结构
	query := "DESCRIBE videos"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("查询表结构失败: %v", err)
	}
	defer rows.Close()

	fmt.Println("videos表结构:")
	fmt.Println("+----------------+------------------+------+-----+-------------------+-------------------+")
	fmt.Println("| Field          | Type             | Null | Key | Default           | Extra             |")
	fmt.Println("+----------------+------------------+------+-----+-------------------+-------------------+")

	var field, type_, null, key, extra string
	var def *string
	for rows.Next() {
		err := rows.Scan(&field, &type_, &null, &key, &def, &extra)
		if err != nil {
			log.Fatalf("扫描行失败: %v", err)
		}
		defStr := "NULL"
		if def != nil {
			defStr = *def
		}
		fmt.Printf("| %-14s | %-16s | %-4s | %-3s | %-17s | %-17s |\n", field, type_, null, key, defStr, extra)
	}
	fmt.Println("+----------------+------------------+------+-----+-------------------+-------------------+")

	// 检查数据库名称
	dbNameQuery := "SELECT DATABASE()"
	var dbName string
	err = db.QueryRow(dbNameQuery).Scan(&dbName)
	if err != nil {
		log.Fatalf("查询数据库名称失败: %v", err)
	}
	fmt.Printf("当前连接的数据库: %s\n", dbName)
}