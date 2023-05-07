package service

import (
	"fmt"
	"tommy-backend/models"

	"github.com/gin-gonic/gin"
)

// editMenu
// @Tags updateMenu
// @param id formData string true "id" default:"none"
// @param navigation formData string true "navigation" default:"tommy222"
// @param image formData string true "image" default:"none"
// @Success 200 {string} json{"code","message"}
// @Router /api/edit-menu [post]
func CreateWork(c *gin.Context) {
	work := models.Work{}
	work.Title = c.PostForm("title")
	work.Content = c.PostForm("content")
	work.Tag = c.PostFormArray("tag[]")
	fmt.Println("work", work)
}
