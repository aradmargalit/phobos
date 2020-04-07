package controllers

import (
	"net/http"
	"server/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddQuickAddHandler adds a new activity to the database
func (e *Env) AddQuickAddHandler(c *gin.Context) {
	var quickAdd models.QuickAdd
	if err := c.ShouldBindJSON(&quickAdd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add the owner ID to the activituy
	uid := c.GetInt("user")

	quickAdd.OwnerID = uid
	record, err := e.DB.InsertQuickAdd(quickAdd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, record)
}

// GetQuickAddsHandler returns all the user's quick-adds
func (e *Env) GetQuickAddsHandler(c *gin.Context) {
	// Pull user out of context to figure out which quick-adds to grab
	uid := c.GetInt("user")

	qa, err := e.DB.GetQuickAddsByUser(uid)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, qa)
}

// DeleteQuickAddHandler returns all the user's activities
func (e *Env) DeleteQuickAddHandler(c *gin.Context) {
	// Pull user out of context to confirm it's safe to delete the activity
	uid := c.GetInt("user")

	quickAddID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	err = e.DB.DeleteQuickAddByID(uid, quickAddID)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, "Successfully deleted quick-add: "+c.Param("id"))
	return
}
