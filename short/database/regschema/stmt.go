package regschema

import (
	"database/sql"
	"shortener/logs"
)

var RegContext Context

type Context struct {
	DB   *sql.DB
	Stmt map[string]*sql.Stmt
}

func RegisterSchema() {
	regShortStmt()
}

func CreateRegisterContext(db *sql.DB, stmt map[string]*sql.Stmt) {
	RegContext = Context{
		DB:   db,
		Stmt: stmt,
	}
}

func (c *Context) RegisterSchema(name, schema string) {
	if stmt, err := c.DB.Prepare(schema); err != nil {
		logs.Error.Panic(name, err)
	} else {
		c.Stmt[name] = stmt
		logs.Info.Println(name, "STMT seccussed.")
	}
}
