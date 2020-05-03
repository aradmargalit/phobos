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

	metadata := r.Group("/metadata")
	{
		metadata.GET("/activity_types", env.ActivityTypesHandler)
	}
	registerAdminHandlers(r, svc)
	registerStravaHandlers(r, svc)

	// If none of registered routes match, serve the client JS
	r.Use(static.Serve("/", static.LocalFile("../deimos/build", true)))
	r.NoRoute(func(c *gin.Context) {
		c.File("../deimos/build")
	})

	fmt.Println("🚀 🌑 Phobos is ready!")
	r.Run(":8080")
}

func registerAdminHandlers(r *gin.Engine, env *controllers.Env) {
	admin := r.Group("/admin")
	admin.Use(middleware.NonProd)
	{
		admin.GET("/seed", env.SeedHandler)
	}
}

func registerStravaHandlers(r *gin.Engine, env *controllers.Env) {
	r.GET("/public/strava/callback", env.StravaCallbackHandler)
	r.GET("/public/strava/webhook", env.StravaWebookVerificationHandler)
	r.POST("/public/strava/webhook", env.StravaWebHookCatcher)

	strava := r.Group("/strava")
	strava.Use(middleware.AuthRequired)
	{
		strava.GET("/auth", env.StravaLoginHandler)
		strava.GET("/deauth", env.StravaDeauthorizationHandler)
	}
}
