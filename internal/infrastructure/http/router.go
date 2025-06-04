package http

import (
	"github.com/PhipattanachaiDev/golang_api-migration/internal/ports"
	"github.com/gin-gonic/gin"
	swagFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(service ports.UserService) *gin.Engine {
	r := gin.Default()
	r.Use(LoggerMiddleware())

	h := NewHandler(service)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "User API is running!"})
	})
	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swagFiles.Handler))
	
	
	auth := r.Group("/users")
	auth.Use(AuthMiddleware())
	{
		auth.GET("", h.GetUsers)
		auth.GET(":id", h.GetUser)
		auth.PUT(":id", h.UpdateUser)
		auth.DELETE(":id", h.DeleteUser)
	}

	r.POST("/users", h.CreateUser) // Public route

	return r
}

