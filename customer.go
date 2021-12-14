package main

import (
	"database/sql"
	"time"
)

type Customers struct {
	Id         int
	FirstName  string    `db:"first_name"`
	LastName   string    `db:"last_name"`
	BirthDate  time.Time `db:"date_of_birth"`
	Address    string
	Status     int
	Email      string    `db:"email"`
	UserName   string    `db:"user_name"`
	Password   string    `db:"user_password"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	DomisiliID Domicile
}

type Domicile struct {
	Id           sql.NullInt32
	DomicileName string `db:"domisili"`
}

type CustomerDomicile struct {
	Id           int
	FirstName    string `db:"first_name"`
	LastName     string `db:"last_name"`
	DomicileName string `db:"domisili"`
}
