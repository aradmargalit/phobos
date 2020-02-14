package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CurrentUserHandler returns the email of the current user
func (e *Env) CurrentUserHandler(c *gin.Context) {
	session := sessions.Default(c)
	uid := session.Get(UserID)
	

	if email, ok := uid.(string); ok {
		u, err := e.DB.GetUserByEmail(string(email))
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"user": u})

	} else {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("could not get email from user cookie"))
	}
}
