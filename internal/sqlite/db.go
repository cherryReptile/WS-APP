package sqlite

import (
	_ "embed"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"os"
)

//go:embed default.sql
var schema string

func GetDb(userUUID string) (db *sqlx.DB, err error) {
	db, err = sqlx.Open("sqlite3", "./storage/users/"+userUUID+"/"+userUUID+".sqlite3")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}

func SetDefaultSchema(db *sqlx.DB) (err error) {
	_, err = db.Exec(schema)
	if err != nil {
		return err
	}

	return err
}

func Create(userUUID string) error {
	path := "./storage/users/" + userUUID
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	db, err := GetDb(userUUID)
	if err != nil {
		return err
	}

	err = SetDefaultSchema(db)
	if err != nil {
		return err
	}
	logrus.Info("default schema staged successfully")

	return nil
}
