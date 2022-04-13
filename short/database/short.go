package database

import (
	"shortener/common"
	"shortener/logs"
)

func NewShortWithDB() *ShortStructWithDB {
	return &ShortStructWithDB{
		context: &DatbaseContext,
	}
}

func (s *ShortStructWithDB) GetShortWithDB(id string) (*ShortStruct, error) {
	funcName := "Get Short SQL."

	tx, err := s.context.DB.Begin()
	if err != nil {
		logs.Warning.Println(funcName, err)
		return nil, err
	}

	row := tx.Stmt(
		s.context.Stmt["GetShort"],
	).QueryRow(id)

	shortStruct := ShortStruct{}
	err = row.Scan(
		&shortStruct.ID,
		&shortStruct.Url,
		&shortStruct.Expires,
	)

	if TxErrorCheckError(funcName, err, tx) {
		return nil, err
	}

	if TxCommit(funcName, tx) {
		return nil, err
	}

	return &shortStruct, nil
}

func (s *ShortStructWithDB) AddShortWithDB(shortParameter *common.ShortParameter, id string) error {
	funcName := "Add Short SQL."

	tx, err := s.context.DB.Begin()
	if err != nil {
		logs.Warning.Println(funcName, err)
		return err
	}

	_, err = tx.Stmt(
		s.context.Stmt["CreateShort"],
	).Exec(
		id,
		shortParameter.Url,
		shortParameter.Expires,
	)

	if TxErrorCheckError(funcName, err, tx) {
		return err
	}

	if TxCommit(funcName, tx) {
		return err
	}

	return nil
}
