package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CurrentUserHandler returns the email of the current user
func (e *Env) CurrentUserHandler(c *gin.Context) {
	uid, ok := c.Get("user")
	if !ok {
		panic("CurrentUserHandler: No user ID")
	}

	u, err := e.DB.GetUserByID(uid.(int))
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"user": u})
}
