package models

import (
	"gorm.io/gorm"
	"time"
)

type IdeaTable struct {
	ID          int    `gorm:"primaryKey;autoIncrement"`
	IdeaName    string `gorm:"type:varchar(255)"`
	Description string
	Author      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

func (IdeaTable) TableName() string {
	return "idea"
}

type UserTable struct {
	ID    int `gorm:"primaryKey;autoIncrement"`
	Name  string
	Color string
}

func (UserTable) TableName() string {
	return "user"
}
