package routes

import (
	"github.com/Hazem-Ben-Abdelhafidh/golangChat/controllers"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controllers.Register)
	}
}
