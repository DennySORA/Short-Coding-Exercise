package database

import (
	"database/sql"
	"shortener/common"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetShortWithDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	referID := "12345678"
	referURL := "http://www.google.com"
	referTime := time.Now().Add(time.Hour)

	prepare := mock.ExpectPrepare(
		"^SELECT$",
	)

	mock.ExpectBegin()
	prepare.ExpectQuery().WillReturnRows(
		sqlmock.NewRows(
			[]string{"id", "url", "expires"},
		).AddRow(
			referID, referURL, referTime,
		))
	mock.ExpectCommit()

	// now we execute our method
	databaseContext := Context{
		DB:   db,
		Stmt: map[string]*sql.Stmt{},
	}

	stmt, err := db.Prepare("SELECT")
	assert.NoError(t, err)
	databaseContext.Stmt["GetShort"] = stmt
	ShortWithDB := ShortStructWithDB{
		context: &databaseContext,
	}

	shortStruct, err := ShortWithDB.GetShortWithDB(referID)
	assert.NoError(t, err)
	assert.Equal(t, referID, shortStruct.ID)
	assert.Equal(t, referURL, shortStruct.Url)
	assert.Equal(t, referTime, shortStruct.Expires)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAddShortWithDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	referID := "12345678"
	referURL := "http://www.google.com"
	referTime := "2023-02-08T09:20:41"

	prepare := mock.ExpectPrepare(
		"^INSERT$",
	)

	mock.ExpectBegin()
	prepare.ExpectExec().WithArgs(
		referID,
		referURL,
		referTime,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// now we execute our method
	databaseContext := Context{
		DB:   db,
		Stmt: map[string]*sql.Stmt{},
	}

	stmt, err := db.Prepare("INSERT")
	assert.NoError(t, err)
	databaseContext.Stmt["CreateShort"] = stmt
	ShortWithDB := ShortStructWithDB{
		context: &databaseContext,
	}
	err = ShortWithDB.AddShortWithDB(
		&common.ShortParameter{
			Url:     referURL,
			Expires: referTime,
		},
		referID,
	)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
