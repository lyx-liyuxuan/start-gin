package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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
		//DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		//DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		//DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		//SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		SkipDefaultTransaction: false, // 是否跳过默认事务
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "gva_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: false,  // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用建表的时候是否自动创建物理外键约束（推荐使用逻辑外键）
	})
	fmt.Println(db, err)

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	//type User struct {
	//	Name string
	//}
	//
	//type UserTwo struct {
	//	Name string
	//}

	//M := db.Migrator()
	//M.CreateTable(&User{}) // 建表
	//fmt.Println(M.HasTable(&User{})) // 判断表是否存在 true
	//fmt.Println(M.DropTable(&User{})) // 删除表 nil

	/*if M.HasTable(&User{}) {
		M.DropTable(&User{})
	} else {
		M.CreateTable(&User{})
	}*/

	/*if M.HasTable(&User{}) {
		M.RenameTable(&User{}, &UserTwo{})
	} else {
		M.RenameTable(&UserTwo{}, &User{})
	}*/
	GLOBAL_DB = db
	TestUserCreate()
}
