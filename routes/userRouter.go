package routes

import (
	controller "github.com/Frhnmj2004/golang-jwt-project/controllers"
	"github.com/Frhnmj2004/golang-jwt-project/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
