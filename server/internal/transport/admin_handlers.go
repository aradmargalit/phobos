package transport

import (
	"net/http"
	"server/internal/middleware"
	"server/internal/service"

	"github.com/gin-gonic/gin"
)

// Aggregator function to make it easy for the router builder to hit these all at once
func registerAdminHandlers(r *gin.Engine, svc service.PhobosAPI) {
	admin := r.Group("/admin")
	admin.Use(middleware.NonProd)
	{
		admin.GET("/seed", makeSeedHandler(svc))
	}
}

func makeSeedHandler(svc service.PhobosAPI) func(*gin.Context) {
	return func(c *gin.Context) {
		err := svc.SeedActivityTypes()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.String(http.StatusOK, "Successfully seeded database. 🌱")
	}
}
