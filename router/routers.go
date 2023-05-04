package routers

import (
	"tommy-backend/service"

	"github.com/gin-gonic/gin"
)

func ApiRouters(r *gin.Engine) {
	apiRouters := r.Group("/api")

	{
		apiRouters.GET("/", service.GetIndex)
		apiRouters.POST("/create-menu", service.CreateMenu)
		apiRouters.GET("/get-menu", service.GetAllMenu)
		apiRouters.POST("/edit-menu", service.EditMenu)
		apiRouters.GET("/create-work")
	}
}
