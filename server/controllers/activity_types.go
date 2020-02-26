package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ActivityTypesHandler returns all available activity types
func (e *Env) ActivityTypesHandler(c *gin.Context) {

	at, err := e.DB.GetActivityTypes()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("could not fetch activity types"))
	}

	c.JSON(http.StatusOK, gin.H{"activity_types": at})
}