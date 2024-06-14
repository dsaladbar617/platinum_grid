package models

import (
	"database/sql"

	"github.com/lib/pq"
)

type User struct {
	ID        int
	Name      string
	ShortName string
	Templates string
}

type UserModel struct {
	DB *sql.DB
}

func (u *UserModel) addUser(firebase_uuid, name, picture, email string) (sql.Result, error) {

	// err := json.NewDecoder(r.Body).Decode(&u)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// resp := make(map[string]string)

	stmt := `INSERT INTO users (firebase_uuid, name, picture, email) VALUES ($1, $2, $3, $4) RETURNING *;`

	result, err := u.DB.Exec(stmt, firebase_uuid, name, picture, email)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				return nil, pgErr
			}
		}

		// http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	return result, nil

}
