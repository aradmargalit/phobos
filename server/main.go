package main

import (
	"net/http"
	controllers "server/controllers"

	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Session Management
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("phobos-auth", store))

	// Called by the UI when the user clicks the "Login with Google Button"
	r.GET("/auth/google", controllers.HandleLogin)

	// Called by Google API once authenticaton is complete
	r.GET("/callback", controllers.HandleCallback)

	private := r.Group("/private")
	private.Use(AuthRequired)
	{
		private.GET("/users/current", controllers.CurrentUserHandler)
	}

	fmt.Println("🚀 Phobos is ready! 🌑")
	r.Run(":8080")
}

// AuthRequired is a simple middleware to check the session
func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(controllers.UserID)
	if user == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}
