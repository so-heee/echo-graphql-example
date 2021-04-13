package server

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/volatiletech/sqlboiler/boil"
)

func (s *Server) OpenDBConnection() *sql.DB {

	dsn := "user:password@tcp(localhost:3306)/sample?charset=utf8mb4&parseTime=True&loc=Local"

	var db *sql.DB
	var err error

	// try connecting to db for several times
	for connectionCount := 1; connectionCount < 5; connectionCount++ {
		db, err = sql.Open("mysql", dsn)

		if err := db.Ping(); err != nil {
			fmt.Printf("can't connect to database: %v. reconnecting...\n\n", err)
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}

	if err != nil {
		fmt.Printf("%v", err)
		panic("failed to connect database")
	}

	if env, found := os.LookupEnv("APP_ENV"); !found || env != "production" {
		boil.DebugMode = true
	}

	return db
}
