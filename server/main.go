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
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/auth/google", auth.HandleLogin)
	r.GET("/callback", auth.HandleCallback)
	r.GET("/currentUser", func(c *gin.Context) {
		v := sessions.Default(c).Get("token")
		fmt.Println(v)
	})
	r.Run(":8080")
}
