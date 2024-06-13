package main

import (
	"fmt"
	"platinum_grid/internal/models"
	"platinum_grid/internal/server"
)

type application struct {
	sheets *models.SheetModel
}

func main() {

	server := server.NewServer()

	fmt.Println("Starting server...")
	err := server.ListenAndServe()

	// app := &application{
	// 	users: &models.SheetModel{DB: dbInstance,}
	// }

	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
