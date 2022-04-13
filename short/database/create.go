package database

import (
	"database/sql"
	"io/ioutil"
	"shortener/database/regschema"
	"shortener/logs"

	"github.com/spf13/viper"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDatabaseContext() func() error {
	db, err := sql.Open(
		"sqlite3",
		viper.GetString("DATABASE_NAME"),
	)
	if err != nil {
		panic("Fatal error database: " + err.Error())
	}

	// check database connection.
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// create database table.
	err = createSQLTable(db)
	if err != nil {
		panic(err)
	}

	// create database context.
	DatbaseContext = Context{
		DB:   db,
		Stmt: map[string]*sql.Stmt{},
	}

	regschema.CreateRegisterContext(
		DatbaseContext.DB,
		DatbaseContext.Stmt,
	)
	regschema.RegisterSchema()
	return db.Close
}

func createSQLTable(db *sql.DB) error {
	path := viper.GetString("CREATE_TABLE_SQL_PATH")
	data, err := ioutil.ReadFile(path)
	if err != nil {
		logs.Error.Panic(err)
	}

	// Create database Tx.
	tx, err := db.Begin()
	if err != nil {
		logs.Error.Panic(err)
	}

	_, err = tx.Exec(string(data))
	if err != nil {
		tx.Rollback()
		logs.Error.Panic(err)
	}

	err = tx.Commit()
	if err != nil {
		logs.Error.Panic(err)
	}

	return nil
}
