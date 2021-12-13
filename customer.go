package main

import "time"

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
