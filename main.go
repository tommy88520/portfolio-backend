package main

import (
	routers "tommy-backend/router"
	"tommy-backend/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(cors.Default())
	routers.ApiRouters(router)
	routers.SwaggerInfo(router)
	utils.InitConfig()
	utils.InitMySQL()
	router.Run(":8080")
}
