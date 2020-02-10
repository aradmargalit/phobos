package main

import (
	auth "server/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", auth.HandleHome)
	r.GET("/login", auth.HandleLogin)
	r.GET("/callback", auth.HandleCallback)
	r.Run(":8080")
}
