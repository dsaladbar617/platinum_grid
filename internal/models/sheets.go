package models

import (
	"database/sql"
	"fmt"
)

type Sheet struct {
	ID        int
	Name      string
	ShortName string
	Templates string
}

type SheetModel struct {
	DB *sql.DB
}

func (m *SheetModel) Insert(name, shortName, templates string) error {

	stmt := `INSERT INTO sheets (name, shortName, templates) VALUES ($1, $2, $3) RETURNING *;`

	_, err := m.DB.Exec(stmt, name, shortName, templates)
	if err != nil {
		fmt.Errorf("Error: %s", err)
		return err
	}

	return nil
}
