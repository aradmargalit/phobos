package transport

import (
	"server/internal/service"

	"github.com/gin-gonic/gin"
)

// alias to quickly say that a function will create a gin handler
type ginHandler func(*gin.Context)

// BuildRouter creates a router with all handlers initialized
func BuildRouter(r *gin.Engine, svc service.PhobosAPI) *gin.Engine {
	registerGoogleAuthHandlers(r, svc)
	registerActivityHandlers(r, svc)
	registerStatisticsHandlers(r, svc)
	registerQuickAddHandlers(r, svc)
	registerMetadataHandlers(r, svc)
	registerAdminHandlers(r, svc)
	registerStravaHandlers(r, svc)

	return r
}
