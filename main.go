package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// same as
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// router.Use(cors.New(config))
	router.Use(cors.Default())
	router.Run()
}
