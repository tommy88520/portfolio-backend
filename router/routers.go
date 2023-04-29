package routers

import (
	"tommy-backend/service"

	"github.com/gin-gonic/gin"
)

func ApiRouters(r *gin.Engine) {
	apiRouters := r.Group("/index")

	{
		apiRouters.GET("/", service.GetIndex)
	}
}
