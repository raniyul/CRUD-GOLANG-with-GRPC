package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addCustomer = `insert into mst_customer(name,customer_type,mobile,email,gender,callback_phone,
		status,created_at)values (?,?,?,?,?,?,?,?)`
	selectCustomerByMobile = `select customer_id,name,customer_type,mobile,email,gender,callback_phone,
		status from mst_customer where mobile = ?`
	selectCustomer = `select customer_id,name,customer_type,mobile,email,gender,callback_phone, status 
		from mst_customer`
	updateCustomer = `update mst_customer set name=?,customer_type=?,mobile=?,email=?,gender=?,
		callback_phone=?,status=? where customer_id=?`
	selectCustomerByEmail = `select customer_id,name,customer_type,mobile,email,gender,callback_phone,
		status from mst_customer where email=?`
)

type dbReadWriter struct {
	db *sql.DB
}

func NewDBReadWriter(url string, schema string, user string, password string) ReadWriter {
	schemaURL := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, url, schema)
	db, err := sql.Open("mysql", schemaURL)
	if err != nil {
		panic(err)
	}
	return &dbReadWriter{db: db}
}

func (rw *dbReadWriter) AddCustomer(customer Customer) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addCustomer, customer.Name, customer.CustomerType, customer.Mobile, customer.Email,
		customer.Gender, customer.CallbackPhone, OnAdd, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadCustomerByMobile(mobile string) (Customer, error) {
	customer := Customer{Mobile: mobile}
	err := rw.db.QueryRow(selectCustomerByMobile, mobile).Scan(&customer.CustomerId, &customer.Name,
		&customer.CustomerType, &customer.Mobile, &customer.Email, &customer.Gender, &customer.CallbackPhone,
		&customer.Status)

	if err != nil {
		return Customer{}, err
	}

	return customer, nil
}

func (rw *dbReadWriter) ReadCustomer() (Customers, error) {
	customer := Customers{}
	rows, _ := rw.db.Query(selectCustomer)
	defer rows.Close()
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.CustomerId, &c.Name, &c.CustomerType, &c.Mobile, &c.Email, &c.Gender,
			&c.CallbackPhone, &c.Status)
		if err != nil {
			fmt.Println("error query:", err)
			return customer, err
		}
		customer = append(customer, c)
	}
	//fmt.Println("db nya:", customer)
	return customer, nil
}

func (rw *dbReadWriter) UpdateCustomer(cus Customer) error {
	//fmt.Println("update")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateCustomer, cus.Name, cus.CustomerType, cus.Mobile, cus.Email, cus.Gender,
		cus.CallbackPhone, cus.Status, cus.CustomerId)

	//fmt.Println("name:", cus.Name, cus.CustomerId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (rw *dbReadWriter) ReadCustomerByEmail(email string) (Customer, error) {
	customer := Customer{Email: email}
	err := rw.db.QueryRow(selectCustomerByEmail, email).Scan(&customer.CustomerId, &customer.Name,
		&customer.CustomerType, &customer.Mobile, &customer.Email, &customer.Gender, &customer.CallbackPhone,
		&customer.Status)

	//fmt.Println("err db", err)
	if err != nil {
		return Customer{}, err
	}

	return customer, nil
}
