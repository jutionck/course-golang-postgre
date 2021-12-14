/*
Requirement (package main)
1. Ubah konfigurasi sql ke dalam os.env (os.Getenv)
2. Buat file baru / sendiri untuk db connection
3. Buat file baru untuk model customer
4. Buat file util untuk konversi tanggal dari string ke date time (YYYY-MM-DD)
*/

package main

import "log"

func main() {

	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// layoutFormat := "2006-01-02"
	// birthDate := "1999-02-22"
	// myDate, _ := DateParsing(layoutFormat, birthDate)

	// // Insert
	// customers := Customers{
	// 	Id:         4,
	// 	FirstName:  "Suci",
	// 	LastName:   "Pardede",
	// 	BirthDate:  myDate,
	// 	Address:    "Jakarta Barat",
	// 	Status:     1,
	// 	Email:      "contoh@gmail.com",
	// 	UserName:   "suci.pardede",
	// 	Password:   "password",
	// 	DomisiliID: 1,
	// 	CreatedAt:  time.Now(),
	// 	UpdatedAt:  time.Now(),
	// }

	// CreateCustomers(db, customers)

	// Update
	// customerUpdate := Customers{
	// 	Id:         1,
	// 	FirstName:  "Jution",
	// 	LastName:   "Kirana",
	// 	BirthDate:  myDate,
	// 	Address:    "Jakarta Barat",
	// 	Status:     1,
	// 	Email:      "contoh@gmail.com",
	// 	UserName:   "jution.kirana",
	// 	Password:   "password",
	// 	DomisiliID: 1,
	// 	CreatedAt:  time.Now(),
	// 	UpdatedAt:  time.Now(),
	// }
	// UpdateCustomer(customerUpdate)

	// Delete
	// customerId := 1
	// DeletCustomer(customerId)

	// Get All Customer
	// GetCustomers(db)

	// Get customer by id
	// GetCustomerByID(db, 3)

	// GetAllCustomerDomicile
	// GetCustomersDomicile(db)

	// Find customer by domisili
	// FindCustomerByDomicile(db, "%Bandung%")

	// Get customer by id with prepare
	GetCustomerByIDWithPrepare(db, 2)

	/*
		Tugas :
		1. Buat update customer by email => email ada atau tidak ? update : "Email tidak ada"
		2. ALTER TABLE mst_customer -> field: is_actived (int) 1 = Active; 0 = Non Active, buat update Active dan Non Active
		3. Buat simulasi login -> is_actived = 1 | 0 = informasi gagal login ...
		4. Tampilkan jumlah customer berdasarkan domisili
		5. Tampilkan rata-rata umur customer berdasarkan domisili
	*/
}
