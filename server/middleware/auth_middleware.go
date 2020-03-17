package middleware

import (
	"fmt"
	"net/http"
	"server/controllers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthRequired is a simple middleware to check the session
func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	uid := session.Get(controllers.UserID)

	if uid == nil {
		fmt.Println("Aborting")
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// Store the user
	// You can get this in any protected route with c.Get("user")
	c.Set("user", uid)

	// Continue down the chain to handler etc
	c.Next()
}
