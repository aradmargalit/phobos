package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CurrentUserHandler returns the email of the current user
func (e *Env) CurrentUserHandler(c *gin.Context) {
	uid := c.GetInt("user")

	u, err := e.DB.GetUserByID(uid)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"user": u})
}
