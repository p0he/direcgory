package main

import (
	"database/sql"

	"direcgory/config"
	"direcgory/repository"

	_ "github.com/mattn/go-sqlite3" // just register driver.
)

func init() {
	config.Configure()
}

func main() {
	db, err := connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := repository.NewDirectoryRepository(db)
	run(r)
}

func connect() (*sql.DB, error) {
	db, err := sql.Open(config.DriverName, config.DataSourceName)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func run(r *repository.DirectoryRepository) error {
	if err := r.CreateTable(); err != nil {
		return err
	}

	for {
		if err := r.ShowAllDirectories(); err != nil {
			return err
		}

		if err := r.Create(); err != nil {
			return err
		}
	}

	return nil
}