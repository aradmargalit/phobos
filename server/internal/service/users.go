package service

import (
	"server/internal/responsetypes"

	"github.com/gin-gonic/gin"
)

// CurrentUserHandler returns the email of the current user
func (svc *service) GetCurrentUser(c *gin.Context) responsetypes.User {
	uid := c.GetInt("user")

	u, err := svc.db.GetUserByID(uid)
	if err != nil {
		panic(err)
	}

	return u
}
