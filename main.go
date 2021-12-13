package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type Customers struct {
	Id         int
	FirstName  string    `db:"first_name"`
	LastName   string    `db:"last_name"`
	BirthDate  time.Time `db:"date_of_birth"`
	Address    string
	Status     int
	Email      string
	UserName   string    `db:"user_name"`
	Password   string    `db:"user_password"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	DomisiliID int       `db:"domisili_id"`
}

// type newCustomer struct {

// }

func main() {

	// Loading config database / bisa menggunakan environment
	dbHost := "localhost"
	dbPort := "5432"
	dbName := "gold_market_db"
	dbUser := "postgres"
	dbPassword := ""

	// Database connection string
	// Cara 1. postgres://user:password@host:port/db?sslmode=disable
	// Cara 2. host= user= dbname= sslmode=disable password= port=
	datasourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open connection to database
	db, err := sqlx.Connect("pgx", datasourceName)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully connect to database!")
	}

	// Close connection to database
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	// Modifying data to database

	// Insert batch / bulk
	// customers := []Customers{
	// 	{
	// 		Id:         1,
	// 		FirstName:  "Maria",
	// 		LastName:   "Kirana",
	// 		BirthDate:  time.Date(1995, 01, 22, 10, 20, 0, 0, time.UTC),
	// 		Address:    "Jakarta Barat",
	// 		Status:     1,
	// 		Email:      "contoh@gmail.com",
	// 		UserName:   "jution.kirana",
	// 		Password:   "password",
	// 		DomisiliID: 1,
	// 		CreatedAt:  time.Now(),
	// 		UpdatedAt:  time.Now()},
	// 	{
	// 		Id:         2,
	// 		FirstName:  "Surya",
	// 		LastName:   "Kirana",
	// 		BirthDate:  time.Date(1995, 01, 22, 10, 20, 0, 0, time.UTC),
	// 		Address:    "Jakarta Barat",
	// 		Status:     1,
	// 		Email:      "contoh@gmail.com",
	// 		UserName:   "jution.kirana",
	// 		Password:   "password",
	// 		DomisiliID: 1,
	// 		CreatedAt:  time.Now(),
	// 		UpdatedAt:  time.Now(),
	// 	},
	// }

	// _, err = db.NamedExec(`INSERT INTO mst_customer
	// 											(id, first_name, last_name, date_of_birth, address, status, email, user_name, user_password, domisili_id, created_at, updated_at)
	// 											VALUES (:id, :first_name, :last_name, :date_of_birth, :address, :status, :email, :user_name, :user_password, :domisili_id, :created_at, :updated_at)`, customers)

	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	log.Println("Successfully insert data to database!")
	// }

	// Insert with map[string]interface{}
	// newCustomer := map[string]interface{}{
	// 	"id":         4,
	// 	"first_name": "Valention",
	// 	"last_name":  "Rosa",
	// 	"status":     1,
	// 	"address":    "Grogol",
	// }

	// _, err = db.NamedExec(`INSERT INTO mst_customer (id, first_name, last_name, status, address)
	// 											VALUES (:id, :first_name, :last_name, :status, :address)`, newCustomer)

	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	log.Println("Successfully insert data to database!")
	// }

	// Update
	// customerUpdate := Customers{
	// 	FirstName:  "Valention",
	// 	LastName:   "Rosa",
	// 	BirthDate:  time.Date(1980, 01, 22, 10, 20, 0, 0, time.UTC),
	// 	Address:    "Grogol",
	// 	Status:     1,
	// 	Email:      "valention.rosa@gmail.com",
	// 	UserName:   "valention.rosa",
	// 	Password:   "password",
	// 	DomisiliID: 1,
	// 	CreatedAt:  time.Now(),
	// 	UpdatedAt:  time.Now(),
	// 	Id:         4,
	// }
	// _, err = db.NamedExec(`UPDATE mst_customer SET first_name=:first_name, last_name=:last_name,
	// 											date_of_birth=:date_of_birth, address=:address, status=:status, email=:email, user_name=:user_name, user_password=:user_password, domisili_id=:domisili_id, created_at=:created_at, updated_at=:updated_at WHERE id=:id`, customerUpdate)

	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	log.Println("Successfully update data to database!")
	// }

	// Delete
	// db.NamedExec("DELETE FROM mst_customer WHER idx =? ")

	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	log.Println("Successfully delete data to database!")
	// }
}

/*
Requirement (package main)
1. Ubah konfigurasi sql ke dalam os.env (os.Getenv)
2. Buat file baru / sendiri untuk db connection
3. Buat file baru untuk model customer
4. Buat file util untuk konversi tanggal dari string ke date time (YYYY-MM-DD)
*/
