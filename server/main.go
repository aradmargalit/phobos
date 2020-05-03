package main

import (
	"server/internal/middleware"
	"server/internal/repository"
	"server/internal/service"

	"fmt"
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func init() {
	for _, v := range []string{
		"GOOGLE_CLIENT_ID",
		"GOOGLE_CLIENT_SECRET",
		"COOKIE_SECRET_TOKEN",
		"FRONTEND_URL",
		"SERVER_URL",
	} {
		if os.Getenv(v) == "" {
			panic(fmt.Sprintf("%v must be set in the environment!", v))
		}
	}
}

func main() {

	// init the database
	db := repository.New()

	// init the service with the database
	svc := service.New(db)

	// Session Management
	// Because this token is random sessions are invalidated when the server restarts
	store := cookie.NewStore([]byte(os.Getenv("COOKIE_SECRET_TOKEN")))
	r.Use(sessions.Sessions("phobos-auth", store))
	// If none of registered routes match, serve the client JS
	r.Use(static.Serve("/", static.LocalFile("../deimos/build", true)))
	r.NoRoute(func(c *gin.Context) {
		c.File("../deimos/build")
	})

	fmt.Println("🚀 🌑 Phobos is ready!")
	r.Run(":8080")
}
