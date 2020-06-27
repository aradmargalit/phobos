package main

import (
	"server/internal/repository"
	"server/internal/service"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	db := repository.New()
	svc := service.New(db)
	start := time.Now()
	svc.BackfillMeters()

	logrus.Info("Took: ", time.Since(start))
}
