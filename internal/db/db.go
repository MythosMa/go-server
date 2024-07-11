package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
	var err error
	dsn := "MythosMa:HakureiReimu16@tcp(127.0.0.1:9001)/go_game_server?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		return fmt.Errorf("数据库连接失败: %v", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("无法连接到数据库: %v", err)
	}

	return nil
}
