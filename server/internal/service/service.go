package service

import "server/internal/models"

type PhobosAPI interface {

}

type service struct{}

//Example

//Could be responsetype
func (svc *service) GetActivities(ctx *gin.Context) ([]models.Activity, error) {
	svc.db.GetActivities()...
}