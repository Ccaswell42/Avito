package storage

import (
	"avito/config"
	"database/sql"
	"log"
)

func ConnectToDb(conf *config.Config) (*sql.DB, error) {

	db, err := sql.Open(conf.DriverName, conf.DSN)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)

	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
}
