package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addMenu = `insert into tbMenu(namaMenu,harga,IDKategoriMenu,createdBy,status,createdOn
		)values (?,?,?,?,?,?)`
	selectMenu = `select IDMenu,namaMenu,harga,IDKategoriMenu,createdBy,status 
		from tbMenu`
	updateMenu = `update tbMenu set namaMenu=?,harga=?,IDKategoriMenu=?,status=?,UpdateBy=?,
		UpdateOn=? where IDMenu=?`
	selectNamaMenu = `select IDMenu,namaMenu,harga,IDKategoriMenu,createdBy,status 
		from tbMenu where namaMenu=?`
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
func (rw *dbReadWriter) AddMenu(menu Menu) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addMenu, menu.NamaMenu, menu.Harga, menu.IDKategoriMenu, menu.CreatedBy, OnAdd, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadMenu() (Menus, error) {
	fmt.Println("read all berhasil")
	menu := Menus{}
	rows, _ := rw.db.Query(selectMenu)
	defer rows.Close()
	for rows.Next() {
		var m Menu
		err := rows.Scan(&m.IDMenu, &m.NamaMenu, &m.Harga, &m.IDKategoriMenu, &m.CreatedBy, &m.Status)
		if err != nil {
			fmt.Println("error query:", err)
			return menu, err
		}
		menu = append(menu, m)
	}
	return menu, nil
}

func (rw *dbReadWriter) UpdateMenu(m Menu) error {
	fmt.Println("update berhasil")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateMenu, m.NamaMenu, m.Harga, m.IDKategoriMenu, m.Status, m.UpdateBy, time.Now(), m.IDMenu)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (rw *dbReadWriter) ReadMenuByNamaMenu(namamenu string) (Menu, error) {
	m := Menu{NamaMenu: namamenu}
	err := rw.db.QueryRow(selectNamaMenu, namamenu).Scan(&m.IDMenu, &m.NamaMenu, &m.Harga, &m.IDKategoriMenu, &m.CreatedBy, &m.Status)

	if err != nil {
		return Menu{}, err
	}

	return m, nil
}
