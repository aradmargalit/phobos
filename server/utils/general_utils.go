package utils

import (
	"time"
)

// Retry will retry the func "fn" as many times as requested
func Retry(fn func() error, tries int, duration time.Duration) (err error) {
	for i := 0; i < tries; i++ {
		if err = fn(); err == nil {
			return
		}
		time.Sleep(duration)
	}
	return
}
