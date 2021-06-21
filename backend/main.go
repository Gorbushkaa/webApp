package main

import (
	
	"webApp/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	router.Use(cors.Default())

	initializeRoutes()

	db.InitialMigration()

	router.Run(":3000")
}
