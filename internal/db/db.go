package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

)

var DB *sql.DB

func InitDB() error {
	// 加载 .env 文件中的环境变量
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")

	// dsn := "MythosMa:HakureiReimu16@tcp(127.0.0.1:3306)/go_game_server?charset=utf8mb4&parseTime=True&loc=Local"
	// 构建 DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	var db *sql.DB
	var err error

	// 重试连接数据库
	for i := 0; i < 10; i++ {
		log.Printf("正在尝试连接到MySQL数据库: %v", dsn)
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Printf("Failed to connect to MySQL (attempt %d): %v", i+1, err)
			time.Sleep(2 * time.Second)
			continue
		}
		err = db.Ping()
		if err == nil {
			return nil
		}
		log.Printf("Failed to ping MySQL (attempt %d): %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	// var err error
	// dsn := "MythosMa:HakureiReimu16@tcp(127.0.0.1:9001)/go_game_server?charset=utf8mb4&parseTime=True&loc=Local"
	// DB, err = sql.Open("mysql", dsn)

	if err != nil {
		return fmt.Errorf("数据库连接失败: %v", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("无法连接到数据库: %v", err)
	}

	return nil
}
