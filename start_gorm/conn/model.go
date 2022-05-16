package main

import (
	"database/sql"
	"time"
)

type TestUser struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivedAt    sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func TestUserCreate() {
	GLOBAL_DB.AutoMigrate(&TestUser{})
}
