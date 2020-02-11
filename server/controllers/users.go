package controllers

import (
	"fmt"
	"net/http"
	models "server/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CurrentUserHandler returns the email of the current user
func CurrentUserHandler(c *gin.Context) {
	session := sessions.Default(c)
	uid := session.Get(UserID)

	if email, ok := uid.(string); ok {
		// Get User from the DB
		db := models.DB{}
		db.Connect()

		u, err := db.GetUserByEmail(string(email))
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"user": u})

	} else {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("could not get email from user cookie"))
	}
}
