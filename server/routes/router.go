package routes

import (
	"booksAPiProject/controllers"
	"booksAPiProject/goi18n"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	goi18n.InitMessages()
	main := router.Group("api/v1")
	{
		health := main.Group("health")
		{
			health.GET("/", controllers.IsHealth)
		}
		books := main.Group("books")
		{
			books.GET("/:id", controllers.GetById())
			books.POST("/", controllers.CreateBook())
			books.GET("/", controllers.GetAll())
			books.PUT("/", controllers.UpdateBook())
			books.DELETE("/:id", controllers.DeleteBook())
		}
	}
	return router
}
