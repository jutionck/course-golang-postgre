package main

const (
	INSERT_CUSTOMERS = `INSERT INTO mst_customer
	(id, first_name, last_name, date_of_birth, address, status, email, user_name, user_password, domisili_id, created_at, updated_at)
	VALUES (:id, :first_name, :last_name, :date_of_birth, :address, :status, :email, :user_name, :user_password, :domisili_id, :created_at, :updated_at)`

	UPDATE_CUSTOMER = `UPDATE mst_customer SET first_name=:first_name, last_name=:last_name,
	date_of_birth=:date_of_birth, address=:address, status=:status, email=:email, user_name=:user_name, user_password=:user_password, domisili_id=:domisili_id, created_at=:created_at, updated_at=:updated_at WHERE id=:id`

	DELETE_CUSTOMER = `DELETE FROM mst_customer WHERE id = $1`
)
