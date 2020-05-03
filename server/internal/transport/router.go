package transport

import (
	"os"
	"server/internal/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// alias to quickly say that a function will create a gin handler
type ginHandler func(*gin.Context)

// BuildRouter creates a router with all handlers initialized
func BuildRouter(svc service.PhobosAPI) *gin.Engine {
	r := gin.Default()

	// Session Management
	// Because this token is random sessions are invalidated when the server restarts
	store := cookie.NewStore([]byte(os.Getenv("COOKIE_SECRET_TOKEN")))
	r.Use(sessions.Sessions("phobos-auth", store))

	registerGoogleAuthHandlers(r, svc)
	registerActivityHandlers(r, svc)

	return r
}
