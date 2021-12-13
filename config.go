package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func InitDB() (*sqlx.DB, error) {
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("DBNAME")
	password := os.Getenv("PASSWORD")

	datasourceName := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)

	// Open connection to database
	db, err := sqlx.Connect("pgx", datasourceName)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully connect to database!")
	}

	// // Close connection to database
	// defer func(db *sqlx.DB) {
	// 	err := db.Close()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }(db)

	return db, nil
}
