package main

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func CreateCustomers(db *sqlx.DB, customers Customers) {
	_, err := db.NamedExec(INSERT_CUSTOMERS, customers)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully insert data to database!")
	}
}

func UpdateCustomer(db *sqlx.DB, customer Customers) {
	_, err := db.NamedExec(UPDATE_CUSTOMER, customer)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully update data to database!")
	}
}

func DeletCustomer(db *sqlx.DB, id int) {
	db.MustExec(DELETE_CUSTOMER, id)
}
