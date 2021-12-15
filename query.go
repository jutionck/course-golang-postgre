package main

const (
	INSERT_CUSTOMERS = `INSERT INTO mst_customer
	(id, first_name, last_name, date_of_birth, address, status, email, user_name, user_password, domisili_id, created_at, updated_at)
	VALUES (:id, :first_name, :last_name, :date_of_birth, :address, :status, :email, :user_name, :user_password, :domisili_id, :created_at, :updated_at)`

	UPDATE_CUSTOMER = `UPDATE mst_customer SET first_name=:first_name, last_name=:last_name,
	date_of_birth=:date_of_birth, address=:address, status=:status, email=:email, user_name=:user_name, user_password=:user_password, domisili_id=:domisili_id, created_at=:created_at, updated_at=:updated_at WHERE id=:id`

	DELETE_CUSTOMER = `DELETE FROM mst_customer WHERE id = $1`

	GET_CUSTOMER_BY_ID_PREPARE = `
	SELECT 
	mst_customer.id, mst_customer.first_name, mst_customer.last_name, mst_domisili.domisili
	FROM
	mst_customer JOIN mst_domisili
	ON mst_customer.domisili_id = mst_domisili.id
	WHERE mst_customer.id = :id
	ORDER BY mst_customer.id`

	UPDATE_SALDO_CUSTOMER_PLUS = `UPDATE mst_customer SET money = (money + :money ) WHERE id = :id`

	UPDATE_SALDO_CUSTOMER_MINUS = `UPDATE mst_customer SET money = (money - :money ) WHERE id = :id`
)
