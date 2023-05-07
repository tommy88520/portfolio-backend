package models

import (
	"tommy-backend/utils"

	"gorm.io/gorm"
)

type Work struct {
	gorm.Model
	Title   string
	Content string
	Tag     []string
}

func (table *Work) TableName() string {
	return "work"
}

func CreateWork(work *Work) *gorm.DB {
	result := utils.DB.Create(&work)
	return result
}
