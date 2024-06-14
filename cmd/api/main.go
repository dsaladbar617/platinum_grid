package main

import (
	"context"
	"fmt"

	// "log"
	"platinum_grid/internal/models"
	"platinum_grid/internal/server"
	// firebase "firebase.google.com/go"
	// "google.golang.org/api/option"
)

type application struct {
	sheets *models.SheetModel
}

const (
	firebaseConfigFile = "/config/smartsheets_firebase.json"
)

var (
	ctx context.Context
)

func main() {
	// ctx = context.Background()
	// opt := option.WithCredentialsFile(firebaseConfigFile)
	// app, err := firebase.NewApp(ctx, nil, opt)
	// if err != nil {
	// 	log.Fatalf("Firebae initialization error: %v\n", err)
	// }

	server := server.NewServer()

	fmt.Println("Starting server...")
	err := server.ListenAndServe()

	// app := &application{
	// 	users: &models.SheetModel{DB: database.Service},
	// }

	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
