package routes

import (
	"Challange-7/handler"
	"Challange-7/service"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, app service.ServiceInterface) {
	server := handler.NewHttpServer(app)
	api := r.Group("/book")
	{
		api.POST("", server.CreateBook)
		api.GET("/all", server.GetBooks)
		api.GET("/:id", server.GetBookById)
		api.PUT("/:id", server.UpdateBook)
		api.DELETE("/:id", server.DeleteBook)
	}
}
