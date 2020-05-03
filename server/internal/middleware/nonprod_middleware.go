package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// NonProd kills requests that shouldn't run in prod
func NonProd(c *gin.Context) {
	if mode := os.Getenv("GIN_MODE"); mode == "release" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "cannot perform this action in release mode! Don't do that"})
	}
}
