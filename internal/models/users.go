package models

import "database/sql"

type User struct {
	ID        int
	Name      string
	ShortName string
	Templates string
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, picture, email string) error {

	return nil
}
