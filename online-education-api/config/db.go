package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// DBConfig 数据库配置
type DBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

// GetDBConfig 从环境变量或配置文件获取数据库配置
func GetDBConfig() *DBConfig {
	// 实际项目中应该从环境变量或配置文件中读取
	// 请根据您的实际数据库配置修改以下信息
	return &DBConfig{
		Host:     "localhost",
		Port:     3306,
		Username: "root",
		Password: "lyd555625", // 默认密码，实际使用时请修改
		DBName:   "online_education_system",
	}
}

// 注意：如果您是首次设置，请确保：
// 1. MySQL服务已启动
// 2. 已创建名为'online_education'的数据库
// 3. 用户名和密码正确
// 4. 用户有访问该数据库的权限

// InitDB 初始化数据库连接
func InitDB(config *DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("无法打开数据库连接: %w", err)
	}

	// 设置连接池参数
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Hour)

	// 测试连接
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("无法连接到数据库: %w", err)
	}

	log.Println("成功连接到数据库")
	return db, nil
}