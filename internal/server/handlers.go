package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/lib/pq"
)

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

type newSheet struct {
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Templates string `json:"templates"`
}

func (s *Server) addSheet(w http.ResponseWriter, r *http.Request) {
	var ns newSheet

	stmt := `INSERT INTO sheets (name, short_name, templates) VALUES ($1, $2, $3) RETURNING *;`

	err := json.NewDecoder(r.Body).Decode(&ns)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = s.db.Exec(stmt, ns.Name, ns.ShortName, ns.Templates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := make(map[string]string)
	resp["message"] = "sheet added successfully"

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)

}

func (s *Server) updateSheet(w http.ResponseWriter, r *http.Request) {
	var ns newSheet

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&ns)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stmt := `UPDATE SHEETS SET
  name = COALESCE(NULLIF($2,''), name),
  short_name = COALESCE(NULLIF($3,''), short_name),
  templates = COALESCE(NULLIF($4,''), templates)
  WHERE id = $1
  RETURNING *;
  `

	_, err = s.db.Exec(stmt, id, ns.Name, ns.ShortName, ns.Templates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := make(map[string]string)
	resp["message"] = "sheet edited successfully"

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)

}

type Sheet struct {
	ID        int32  `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Templates string `json:"templates"`
}

func (s *Server) getSheetByID(w http.ResponseWriter, r *http.Request) {
	var sheet Sheet
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	stmt := `SELECT * FROM sheets WHERE id = $1 LIMIT 1;`

	_ = s.db.QueryRow(stmt, id).Scan(&sheet.ID, &sheet.Name, &sheet.ShortName, &sheet.Templates)

	response, err := json.Marshal(sheet)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(response)

}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

type User struct {
	UserID  string `json:"firebase_uuid"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
	Email   string `json:"email"`
}

func (s *Server) addUser(w http.ResponseWriter, r *http.Request) {
	var u User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := make(map[string]string)

	stmt := `INSERT INTO users (firebase_uuid, name, picture, email) VALUES ($1, $2, $3, $4) RETURNING *;`

	_, err = s.db.Exec(stmt, u.UserID, u.Name, u.Picture, u.Email)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" {
				w.WriteHeader(http.StatusForbidden)
				resp["message"] = "User already exists"
				json.NewEncoder(w).Encode(resp)
				// http.Error(w, err.Error(), http.StatusForbidden)
				return
			}
		}

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp["message"] = fmt.Sprintf("%s added successfully", u.Name)

	json.NewEncoder(w).Encode(resp)

}
