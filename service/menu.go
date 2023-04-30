package service

import (
	"fmt"
	"tommy-backend/models"

	"github.com/gin-gonic/gin"
)

// GetIndex
// @Tags Index
// @Success 200 {string} welcome
// @Router /user [get]
func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welcome",
	})
}

// CreateMenu
// @Tags CreateMenu
// @param navigation formData string true "navigation" default:"tommy222"
// @param image formData string true "image" default:"none"
// @Success 200 {string} json{"code","message"}
// @Router /menu/create-menu [post]
func CreateMenu(c *gin.Context) {
	menu := &models.Menu{}
	// menu.Navigation = c.PostForm("navigation")
	// menu.Image = c.PostForm("image")
	menu.Navigation = "tommy"
	menu.Image = "tommy"
	fmt.Println("menu", menu)
	// _, err := govalidator.ValidateStruct(menu)
	// if err != nil {
	// 	fmt.Println(err)
	// 	c.JSON(200, gin.H{
	// 		"code":    -1,
	// 		"message": "create failed",
	// 	})
	// } else {
	models.CreateMenu(menu)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "Success create",
		"data":    menu,
	})
	// }
}

func GetAllKites(c *gin.Context) {
	result := models.GetMenu()
	c.JSON(200, gin.H{
		"code":    0,
		"message": "Getttt",
		"data":    result,
	})
}
