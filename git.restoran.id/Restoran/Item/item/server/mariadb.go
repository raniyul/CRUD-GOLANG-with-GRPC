package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addItem = `insert into tbItem(namaItem,merk,IDSupplier,status,createdBy,createdOn,keterangan)
		values (?,?,?,?,?,?)`
	selectItemByIDSupplier = `select IDItem,namaItem, merk, IDSupplier, status,createdBy,createdOn,
		updateBy, updateOn from tbItem where IDSupplier = ?`
	selectItemBystatus = `select IDItem, namaItem, merk, IDSupplier, status, createdBy, createdOn,
		updateBy, updateOn from tbItem where status=?`
	selectItem = `select IDItem, namaItem, merk, IDSupplier, status, createdBy, createdOn, 
		updateBy, updateOn from tbItem where status=1`
	updateItem = `update tbItem set namaItem=?, merk=?, IDSupplier=?, status=?,
		updateBy=?, updateOn=? where IDItem=?`
	selectKeterangan = `select IDItem,namaItem,merk,IDSupplier,status,createdBy, keterangan from tbItem
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

func (rw *dbReadWriter) AddItem(item Item) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addItem, item.NamaItem, item.Merk, item.IDSupplier, item.Status,
		item.CreatedBy, time.Now(), item.Keterangan) //masih bingung???????
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadItemByIDSupplier(IDSupplier int32) (Item, error) {
	item := Item{IDSupplier: IDSupplier}
	err := rw.db.QueryRow(selectItemByIDSupplier, IDSupplier).Scan(&item.IDItem, &item.NamaItem,
		&item.Merk, &item.IDSupplier, &item.Status, &item.CreatedBy, &item.CreatedOn, &item.UpdateBy,
		&item.UpdateOn)

	if err != nil {
		return Item{}, err
	}

	return item, nil
}

func (rw *dbReadWriter) ReadItem() (Items, error) {
	item := Items{}
	rows, _ := rw.db.Query(selectItem)
	defer rows.Close()
	for rows.Next() {
		var c Item
		err := rows.Scan(&c.IDItem, &c.NamaItem, &c.Merk, &c.IDSupplier, &c.Status, &c.CreatedBy, &c.CreatedOn,
			&c.UpdateBy, &c.UpdateOn, &c.Keterangan)
		if err != nil {
			fmt.Println("error query:", err)
			return item, err
		}
		item = append(item, c)
	}
	//fmt.Println("db nya:", item)
	return item, nil
}

func (rw *dbReadWriter) UpdateItem(cus Item) error {
	//fmt.Println("update")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateItem, cus.NamaItem, cus.Merk, cus.IDSupplier, &cus.Status,
		cus.UpdateBy, time.Now(), cus.IDItem)

	//fmt.Println("namaItem:", cus.namaItem, cus.IDItem)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (rw *dbReadWriter) ReadItemBystatus(status int32) (Item, error) {
	item := Item{Status: status}
	err := rw.db.QueryRow(selectItemBystatus, status).Scan(&item.IDItem, &item.NamaItem,
		&item.Merk, &item.IDSupplier, &item.Status, &item.CreatedBy, &item.CreatedOn,
		&item.UpdateOn, &item.UpdateBy)

	//fmt.Println("err db", err)
	if err != nil {
		return Item{}, err
	}

	return item, nil
}

func (rw *dbReadWriter) ReadByKeterangan(keterangan string) (Items, error) {
	fmt.Println("read all berhasil")
	supplier := Items{}
	rows, _ := rw.db.Query(selectKeterangan, keterangan)
	defer rows.Close()
	for rows.Next() {
		var m Item
		err := rows.Scan(&m.IDItem, &m.NamaItem, &m.Merk, &m.IDSupplier, &m.Status, &m.CreatedBy, &m.Keterangan)
		if err != nil {
			fmt.Println("error query:", err)
			return supplier, err
		}
		supplier = append(supplier, m)
	}
	return supplier, nil
}
