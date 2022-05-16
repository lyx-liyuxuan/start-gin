package main

import (
	"database/sql"
	"time"
)

type Model struct {
	UUID uint      `gorm:"primaryKey"`
	Time time.Time `gorm:"column:my_time"`
}

type TestUser struct {
	Model        Model   `gorm:"embedded;embeddedPrefix:li_"`
	Name         string  `gorm:"default:li"`
	Email        *string `gorm:"not null"`
	Age          uint8   `gorm:"comment:年龄"`
	Birthday     *time.Time
	MemberNumber sql.NullString
}

func TestUserCreate() {
	GLOBAL_DB.AutoMigrate(&TestUser{})
}
