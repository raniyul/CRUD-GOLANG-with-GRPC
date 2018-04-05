package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addPelanggan = `insert into tbPelanggan(namaPelanggan,telepon,email,status,createdBy,createdOn, keterangan
		)values (?,?,?,?,?,?,?)`
	selectPelanggan = `select IDPelanggan,namaPelanggan,telepon,email,status,createdBy from tbPelanggan`
	updatePelanggan = `update tbPelanggan set namaPelanggan=?,telepon=?,email=?,status=?,UpdateBy=?,
		UpdateOn=? where IDPelanggan=?`
	selectNamaPelanggan = `select IDPelanggan,namaPelanggan,telepon,email,status,createdBy 
		from tbPelanggan where namaPelanggan=?`
	selectKeterangan = `select IDPelanggan,namaPelanggan,telepon,email,status,createdBy, keterangan from tbPelanggan
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
func (rw *dbReadWriter) AddPelanggan(pelanggan Pelanggan) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addPelanggan, pelanggan.NamaPelanggan, pelanggan.Telepon, pelanggan.Email, OnAdd, pelanggan.CreatedBy, time.Now(), pelanggan.Keterangan)
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadPelanggan() (Pelanggans, error) {
	fmt.Println("read all berhasil")
	pelanggan := Pelanggans{}
	rows, _ := rw.db.Query(selectPelanggan)
	defer rows.Close()
	for rows.Next() {
		var m Pelanggan
		err := rows.Scan(&m.IDPelanggan, &m.NamaPelanggan, &m.Telepon, &m.Email, &m.Status, &m.CreatedBy)
		if err != nil {
			fmt.Println("error query:", err)
			return pelanggan, err
		}
		pelanggan = append(pelanggan, m)
	}
	return pelanggan, nil
}

func (rw *dbReadWriter) UpdatePelanggan(m Pelanggan) error {
	fmt.Println("update berhasil")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updatePelanggan, m.NamaPelanggan, m.Telepon, m.Email, m.Status, m.UpdateBy, time.Now(), m.IDPelanggan)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (rw *dbReadWriter) ReadByNamaPelanggan(namapelanggan string) (Pelanggan, error) {
	m := Pelanggan{NamaPelanggan: namapelanggan}
	err := rw.db.QueryRow(selectNamaPelanggan, namapelanggan).Scan(&m.IDPelanggan, &m.NamaPelanggan, &m.Telepon, &m.Email, &m.Status, &m.CreatedBy)

	if err != nil {
		return Pelanggan{}, err
	}

	return m, nil
}

func (rw *dbReadWriter) ReadByKeterangan(keterangan string) (Pelanggans, error) {
	fmt.Println("read all berhasil")
	pelanggan := Pelanggans{}
	rows, _ := rw.db.Query(selectKeterangan, keterangan)
	defer rows.Close()
	for rows.Next() {
		var m Pelanggan
		err := rows.Scan(&m.IDPelanggan, &m.NamaPelanggan, &m.Telepon, &m.Email, &m.Status, &m.CreatedBy, &m.Keterangan)
		if err != nil {
			fmt.Println("error query:", err)
			return pelanggan, err
		}
		pelanggan = append(pelanggan, m)
	}
	return pelanggan, nil
}
