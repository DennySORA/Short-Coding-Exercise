package database

import (
	"database/sql"
	"time"
)

var DatbaseContext Context

type Context struct {
	DB   *sql.DB
	Stmt map[string]*sql.Stmt
}

type ShortStruct struct {
	ID      string
	Url     string
	Expires time.Time
}

type ShortStructWithDB struct {
	context *Context
}
