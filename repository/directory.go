package repository

import (
	"database/sql"
	"fmt"
)

type DirectoryRepository struct {
	db *sql.DB
}

func NewDirectoryRepository(db *sql.DB) *DirectoryRepository {
	return &DirectoryRepository{db: db}
}

func (r *DirectoryRepository) CreateTable() error {
	const sql = `
	CREATE TABLE IF NOT EXISTS directory (
			id    INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name  TEXT NOT NULL,
			phone TEXT NOT NULL
	);`

	if _, err := r.db.Exec(sql); err != nil {
		return err
	}

	return nil
}

func (r *DirectoryRepository) ShowAllDirectories() error {
	fmt.Println("Show All Directories")
	rows, err := r.db.Query("SELECT * FROM directory")
	if err != nil {
		return err
	}
	for rows.Next() {
		var d Directory
		if err := rows.Scan(&d.ID, &d.Name, &d.Phone); err != nil {
			return err
		}
		fmt.Printf("[%d] Name:%s TEL:%s\n", d.ID, d.Name, d.Phone)
	}
	fmt.Println("--------")
	return nil
}

func (r *DirectoryRepository) Create() error {
	var d Directory

	fmt.Print("Name >")
	fmt.Scan(&d.Name)

	fmt.Print("TEL >")
	fmt.Scan(&d.Phone)

	const sql = "INSERT INTO directory(name, phone) values (?. ?)"

	stmt, err := r.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = r.db.Exec(d.Name, d.Phone)
	if err != nil {
		return err
	}
	return nil
}