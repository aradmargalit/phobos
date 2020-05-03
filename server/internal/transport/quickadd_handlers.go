package transport

import (
	"server/internal/middleware"
	"server/internal/service"

	"github.com/gin-gonic/gin"
)

// Aggregator function to make it easy for the router builder to hit these all at once
func registerQuickAddHandlers(r *gin.Engine, svc service.PhobosAPI) {
	private := r.Group("/private")
	private.Use(middleware.AuthRequired)

	{
		private.GET("/quick_adds", env.GetQuickAddsHandler)
		private.POST("/quick_adds", env.AddQuickAddHandler)
		private.DELETE("/quick_adds/:id", env.DeleteQuickAddHandler)
	}
}
