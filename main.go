package main

import (
	"context"
	"os"
	"github.com/ninoamine/A-S-Controllers---Postgresql/pkg/postgresql"
)


func main() {
	url := os.Getenv("DATABASE_URL")
	conn, err := postgresql.ConnectToDB(url)
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close(context.Background())
}