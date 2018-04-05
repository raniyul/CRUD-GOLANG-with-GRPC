package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addSupplier = `insert into tbSupplier(namaSupplier,alamat,telepon,email,status,createdBy,createdOn,keterangan
		)values (?,?,?,?,?,?,?,?)`
	selectSupplier = `select IDSupplier,namaSupplier,alamat,telepon,email,status,createdBy from tbSupplier`
	updateSupplier = `update tbSupplier set namaSupplier=?,alamat=?,telepon=?,email=?,status=?,UpdateBy=?,
		UpdateOn=? where IDSupplier=?`
	selectNamaSupplier = `select IDSupplier,namaSupplier,alamat,telepon,email,status,createdBy 
		from tbSupplier where namaSupplier=?`
	selectKeterangan = `select IDSupplier,namaSupplier,alamat,telepon,email,status,createdBy, keterangan from tbSupplier
	 	where keterangan LIKE ?`
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
func (rw *dbReadWriter) AddSupplier(supplier Supplier) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addSupplier, supplier.NamaSupplier, supplier.Alamat, supplier.Telepon, supplier.Email, OnAdd, supplier.CreatedBy, time.Now(), supplier.Keterangan)
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadSupplier() (Suppliers, error) {
	fmt.Println("read all berhasil")
	supplier := Suppliers{}
	rows, _ := rw.db.Query(selectSupplier)
	defer rows.Close()
	for rows.Next() {
		var m Supplier
		err := rows.Scan(&m.IDSupplier, &m.NamaSupplier, &m.Alamat, &m.Telepon, &m.Email, &m.Status, &m.CreatedBy, &m.Keterangan)
		if err != nil {
			fmt.Println("error query:", err)
			return supplier, err
		}
		supplier = append(supplier, m)
	}
	return supplier, nil
}

func (rw *dbReadWriter) UpdateSupplier(m Supplier) error {
	fmt.Println("update berhasil")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateSupplier, m.NamaSupplier, m.Alamat, m.Telepon, m.Email, m.Status, m.UpdateBy, time.Now(), m.IDSupplier)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (rw *dbReadWriter) ReadByNamaSupplier(namasupplier string) (Supplier, error) {
	m := Supplier{NamaSupplier: namasupplier}
	err := rw.db.QueryRow(selectNamaSupplier, namasupplier).Scan(&m.IDSupplier, &m.NamaSupplier, &m.Alamat, &m.Telepon, &m.Email, &m.Status, &m.CreatedBy)

	if err != nil {
		return Supplier{}, err
	}

	return m, nil
}

func (rw *dbReadWriter) ReadByKeterangan(keterangan string) (Suppliers, error) {
	fmt.Println("read all berhasil")
	supplier := Suppliers{}
	rows, _ := rw.db.Query(selectKeterangan, keterangan)
	defer rows.Close()
	for rows.Next() {
		var m Supplier
		err := rows.Scan(&m.IDSupplier, &m.NamaSupplier, &m.Alamat, &m.Telepon, &m.Email, &m.Status, &m.CreatedBy, &m.Keterangan)
		if err != nil {
			fmt.Println("error query:", err)
			return supplier, err
		}
		supplier = append(supplier, m)
	}
	return supplier, nil
}
