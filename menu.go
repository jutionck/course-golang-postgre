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

func GetCustomers(db *sqlx.DB) {
	customers := []Customers{}
	db.Select(&customers, `SELECT id,first_name,last_name,date_of_birth,email FROM mst_customer order by id`)

	// kirana, johan, budi := customers[0], customers[1], customers[2]
	// log.Printf("%#v\n%#v\n%#v\n", kirana, johan, budi)
	log.Println(len(customers))
}

func GetCustomerByID(db *sqlx.DB, id int) {
	customer := Customers{}
	err := db.Get(&customer, `SELECT id,first_name,last_name,date_of_birth,email FROM mst_customer WHERE id = $1 order by id`, id)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(customer)
}

func GetCustomersDomicile(db *sqlx.DB) {
	var customers []CustomerDomicile
	query := `
	SELECT 
	mst_customer.id, mst_customer.first_name, mst_customer.last_name, mst_domisili.domisili
	FROM
	mst_customer JOIN mst_domisili
	ON mst_customer.domisili_id = mst_domisili.id
	ORDER BY mst_customer.id`

	rows, err := db.Queryx(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		customer := CustomerDomicile{}
		err := rows.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.DomicileName)
		if err != nil {
			log.Fatal(err)
		}

		customers = append(customers, customer)
	}

	for _, result := range customers {
		log.Println(result)
	}
}

func FindCustomerByDomicile(db *sqlx.DB, domicile string) {
	var customers []CustomerDomicile
	query := `
	SELECT 
	mst_customer.id, mst_customer.first_name, mst_customer.last_name, mst_domisili.domisili
	FROM
	mst_customer JOIN mst_domisili
	ON mst_customer.domisili_id = mst_domisili.id
	WHERE mst_domisili.domisili ilike $1
	ORDER BY mst_customer.id`

	rows, err := db.Queryx(query, domicile)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		customer := CustomerDomicile{}
		err := rows.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.DomicileName)
		if err != nil {
			log.Fatal(err)
		}

		customers = append(customers, customer)
	}

	for _, result := range customers {
		log.Println(result)
	}
}

// Prepare
func GetCustomerByIDWithPrepare(db *sqlx.DB, id int) {
	stmt, err := db.PrepareNamed(GET_CUSTOMER_BY_ID_PREPARE)
	if err != nil {
		log.Fatal(err)
	}

	customer := CustomerDomicile{}
	customerID := map[string]interface{}{"id": id}

	err = stmt.Get(&customer, customerID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(customer)
}
