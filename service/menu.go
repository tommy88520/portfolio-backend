package service

import (
	"fmt"
	"strconv"
	"tommy-backend/models"
	"tommy-backend/utils"

	"github.com/asaskevich/govalidator"
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
// @Router /api/create-menu [post]
func CreateMenu(c *gin.Context) {
	menu := &models.Menu{}
	menu.Navigation = utils.FirstUpper(c.PostForm("navigation"))
	menu.Image = c.PostForm("image")
	_, err := govalidator.ValidateStruct(menu)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "create failed",
		})
	} else {
		models.CreateMenu(menu)
		c.JSON(200, gin.H{
			"code":    0,
			"message": "Success create",
			"data":    menu,
		})
	}

}

// GetMenu
// @Tags GetMenu
// @Success 200 {string} json{"code","message"}
// @Router /api/get-menu [get]
func GetAllMenu(c *gin.Context) {
	result := models.GetMenu()
	c.JSON(200, gin.H{
		"code":    0,
		"message": "Get menu",
		"data":    result,
	})
}

// updateMenu
// @Tags updateMenu
// @param id formData string true "id" default:"none"
// @param navigation formData string true "navigation" default:"tommy222"
// @param image formData string true "image" default:"none"
// @Success 200 {string} json{"code","message"}
// @Router /api/get-menu [post]
func EditMenu(c *gin.Context) {
	menu := models.Menu{}
	// result := model.EditMenu(c.PostForm(id))
	id, _ := strconv.Atoi(c.PostForm("id"))
	menu.Navigation = utils.FirstUpper(c.PostForm("navigation"))
	menu.Image = c.PostForm("image")
	_, err := govalidator.ValidateStruct(menu)

	if err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "edit failed",
			"data":    menu,
		})
	} else {
		models.EditMenu(id, menu)
		c.JSON(200, gin.H{
			"code":    0,
			"message": "Success edit",
			"data":    menu,
		})
	}

}
