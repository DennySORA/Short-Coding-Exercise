package database

import (
	"database/sql"
	"runtime/debug"
	"shortener/logs"

	"github.com/spf13/viper"
)

func TxErrorCheckError(name string, err error, tx *sql.Tx) bool {
	if err != nil {
		if viper.GetBool("LOG_TX_STATS") {
			logs.Warning.Println(name, err, string(debug.Stack()), err)
		} else {
			logs.Warning.Println(name, err)
		}
		if err := tx.Rollback(); err != nil {
			logs.Error.Println(err)
		}
		return true
	}
	return false
}

func TxCommit(name string, tx *sql.Tx) bool {
	err := tx.Commit()
	if err != nil {
		if viper.GetBool("LOG_TX_STATS") {
			logs.Warning.Println(name, err, string(debug.Stack()), err)
		} else {
			logs.Warning.Println(name, err)
		}
		return true
	}
	return false
}
