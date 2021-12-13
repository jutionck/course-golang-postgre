package main

import (
	"log"
)

func CreateCustomers(customers Customers) {
	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	_, err = db.NamedExec(INSERT_CUSTOMERS, customers)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully insert data to database!")
	}
}

func UpdateCustomer(customer Customers) {
	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	_, err = db.NamedExec(UPDATE_CUSTOMER, customer)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully update data to database!")
	}
}

func DeletCustomer(id int) {
	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	db.MustExec(DELETE_CUSTOMER, id)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully delete data to database!")
	}
}
