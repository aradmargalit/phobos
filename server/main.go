package main

import (
	controllers "server/controllers"
	middleware "server/middleware"
	"server/utils"

	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Session Management
	// Because this token is random sessions are invalidated when the server restarts
	store := cookie.NewStore([]byte(utils.RandomToken()))
	r.Use(sessions.Sessions("phobos-auth", store))

	// Called by the UI when the user clicks the "Login with Google Button"
	r.GET("/auth/google", controllers.HandleLogin)

	// Called by Google API once authenticaton is complete
	r.GET("/callback", controllers.HandleCallback)

	r.GET("/users/logout", controllers.Logout)

	private := r.Group("/private")
	private.Use(middleware.AuthRequired)
	{
		private.GET("/users/current", controllers.CurrentUserHandler)
	}

	fmt.Println("🚀🌑 Phobos is ready!")
	r.Run(":8080")
}
