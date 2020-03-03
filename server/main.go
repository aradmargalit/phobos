package main

import (
	controllers "server/controllers"
	middleware "server/middleware"
	models "server/models"

	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func init() {
	for _, v := range []string{"GOOGLE_CLIENT_ID", "GOOGLE_CLIENT_SECRET", "COOKIE_SECRET_TOKEN"} {
		if os.Getenv(v) == "" {
			panic(fmt.Sprintf("%v must be set in the environment!", v))
		}
	}
}

func main() {
	r := gin.Default()

	db := models.DB{}
	db.Connect()

	env := &controllers.Env{DB: &db}

	// Session Management
	// Because this token is random sessions are invalidated when the server restarts
	store := cookie.NewStore([]byte(os.Getenv("COOKIE_SECRET_TOKEN")))
	r.Use(sessions.Sessions("phobos-auth", store))

	// CORS to allow localhost in development
	// Make sure to allow credentials so we can read the cookie
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true

	r.Use(cors.New(config))

	// Called by the UI when the user clicks the "Login with Google Button"
	r.GET("/auth/google", env.HandleLogin)

	// Called by Google API once authenticaton is complete
	r.GET("/callback", env.HandleCallback)

	r.GET("/users/logout", env.Logout)

	private := r.Group("/private")
	private.Use(middleware.AuthRequired)
	{
		private.GET("/users/current", env.CurrentUserHandler)
		private.POST("/activities", env.AddActivityHandler)
		private.GET("/activities", env.GetActivitiesHandler)
		private.DELETE("/activities/:id", env.DeleteActivityHandler)
		private.PUT("/activities/:id", env.UpdateActivityHandler)
	}

	metadata := r.Group("/metadata")
	{
		metadata.GET("/activity_types", env.ActivityTypesHandler)
	}

	admin := r.Group("/admin")
	admin.Use(middleware.NonProd)
	{
		admin.GET("/seed", env.SeedHandler)
	}

	fmt.Println("🚀 🌑 Phobos is ready!")
	r.Run(":8080")
}
