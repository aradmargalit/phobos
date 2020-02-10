package main

import (
	auth "server/controllers"

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
	r.GET("/auth/google", auth.HandleLogin)

	// Called by Google API once authenticaton is complete
	r.GET("/callback", auth.HandleCallback)
	r.GET("/currentUser", func(c *gin.Context) {
		v := sessions.Default(c).Get("token")
		fmt.Println(v)
	})
	r.Run(":8080")
}
