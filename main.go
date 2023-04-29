package main

import (
	routers "tommy-backend/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(cors.Default())
	routers.ApiRouters(router)
	routers.SwaggerInfo(router)
	router.Run(":8080")
}
