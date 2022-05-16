package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var GLOBAL_DB *gorm.DB

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_study?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               "root:123456@tcp(127.0.0.1:3306)/gorm_study?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize: 191,                                                                                   // string 类型字段的默认长度
	}), &gorm.Config{
		SkipDefaultTransaction:                   false, // 是否跳过默认事务
		DisableForeignKeyConstraintWhenMigrating: true,  // 禁用建表的时候是否自动创建物理外键约束（推荐使用逻辑外键）
	})
	fmt.Println(db, err)

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	GLOBAL_DB = db
	TestUserCreate()
}
