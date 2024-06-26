package routes

import "github.com/gin-gonic/gin"

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticat())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/user_id", controller.GetUser)
}
