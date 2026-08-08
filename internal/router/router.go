package router

import (
	"github.com/AbhinanKumar/smart-dispatch/internal/handler"
	"github.com/AbhinanKumar/smart-dispatch/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(authHandler *handler.AuthHandler, customerHandler *handler.CustomerHandler) *gin.Engine {
	r := gin.Default()
	r.GET("/health", handler.Health)
	r.POST("/register", authHandler.Register)

	r.POST("/login", authHandler.Login)

	protected := r.Group("/")
	protected.Use(midddleware.AuthMiddleware())
	{
		protected.GET("/profile", authHandler.Profile)

		protected.POST("/customers", customerHandler.Create)
		protected.GET("/customers", customerHandler.GetAll)
		protected.GET("/customers/:id", customerHandler.GetByID)
		protected.PUT("/customers/:id", customerHandler.Update)

		// protected.POST("/requests", requestHandler.Create)
		// protected.GET("/requests", requestHandler.GetAll)
		// protected.GET("/requests/:id", requestHandler.GetByID)
		// protected.PUT("/requests/:id", requestHandler.Update)
	}
	return r
}
