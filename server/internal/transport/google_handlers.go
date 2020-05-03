package transport

import (
	"server/internal/service"

	"github.com/gin-gonic/gin"
)

// Aggregator function to make it easy for the router builder to hit these all at once
func registerGoogleAuthHandlers(r *gin.Engine, svc service.PhobosAPI) {
	// Called by the UI when the user clicks the "Login with Google Button"
	r.GET("/auth/google", makeLoginHandler(svc))

	// Called by Google API once authenticaton is complete
	r.GET("/callback", makeCallBackHandler(svc))

	r.GET("/users/logout", makeLogoutHandler(svc))
}

func makeLoginHandler(svc service.PhobosAPI) func(*gin.Context) {
	return func(c *gin.Context) {
		svc.HandleLogin(c)
	}
}

func makeCallBackHandler(svc service.PhobosAPI) func(*gin.Context) {
	return func(c *gin.Context) {
		svc.HandleCallback(c)
	}
}

func makeLogoutHandler(svc service.PhobosAPI) func(*gin.Context) {
	return func(c *gin.Context) {
		svc.Logout(c)
	}
}
