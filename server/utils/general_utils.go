package utils

import (
	"database/sql"
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

// MakeI64 converts an int to a nullable Int64
func MakeI64(i int) sql.NullInt64 {
	return sql.NullInt64{Int64: int64(i), Valid: true}
}
