/*
Requirement (package main)
1. Ubah konfigurasi sql ke dalam os.env (os.Getenv)
2. Buat file baru / sendiri untuk db connection
3. Buat file baru untuk model customer
4. Buat file util untuk konversi tanggal dari string ke date time (YYYY-MM-DD)
*/

package main

func main() {

	// db, err := InitDB()
	// if err != nil {
	// 	log.Fatal(err)
	// }

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
}
