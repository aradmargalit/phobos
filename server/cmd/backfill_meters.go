package main

import (
	"fmt"
	"server/internal/repository"
	"server/internal/service"
	"time"
)

func main() {
	db := repository.New()
	svc := service.New(db)
	start := time.Now()
	svc.BackfillMeters()

	fmt.Println("Took: ", time.Since(start))
}
