package models

import (
	"tommy-backend/utils"

	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Navigation string
	Image      string
}

func (table *Menu) TableName() string {
	return "menu"
}

func CreateMenu(menu *Menu) *gorm.DB {
	result := utils.DB.Create(&menu)
	return result
}

func GetMenu() []*Menu {
	data := make([]*Menu, 10)
	utils.DB.Find(&data)
	return data
}

func EditMenu(id int, menu Menu) *gorm.DB {
	result := utils.DB.Model(&menu).Where("id=?", id).Updates(Menu{Navigation: menu.Navigation, Image: menu.Image})
	return result
}
