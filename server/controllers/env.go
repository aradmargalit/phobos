package controllers

import "server/models"

// Env provides a global environment for handlers and database connections to share
type Env struct {
	DB *models.DB
}
