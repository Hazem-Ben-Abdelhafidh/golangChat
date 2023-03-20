package main

import (
	"github.com/Hazem-Ben-Abdelhafidh/golangChat/config"
	"github.com/Hazem-Ben-Abdelhafidh/golangChat/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDB()
	routes.UserRouter(r)
	r.Run(":6000")
}
