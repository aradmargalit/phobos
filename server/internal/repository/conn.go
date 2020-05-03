package repository

import (
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"

	// Driver for MySQL
	_ "github.com/go-sql-driver/mysql"
)

// Connect connects to the specified database
// The database may not yet be ready, so we're going to retry every second for 30 seconds
func (db *db) Connect() {
	var err error
	var conn *sqlx.DB

	oneSecond, _ := time.ParseDuration("1s")

	for i := 0; i < 30; i++ {
		conn, err = sqlx.Connect("mysql", os.Getenv("API_DB_STRING"))
		if err != nil {
			fmt.Println("Error! ", err, " retrying another ", (30 - i), " times...")
			time.Sleep(oneSecond)
		} else {
			err = nil
			break
		}
	}

	if err != nil {
		panic(err)
	}

	// store the connection in the db struct once established
	db.conn = conn
}
