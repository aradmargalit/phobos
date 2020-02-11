package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CurrentUserHandler returns the email of the current user
func CurrentUserHandler(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(UserID)
	c.JSON(http.StatusOK, gin.H{"user": user})
}
