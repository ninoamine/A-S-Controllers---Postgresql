package main

import (
	"context"
	"os"
	"testing"
	"github.com/jackc/pgx/v5"
)


func TestPostgresqlConnection(t *testing.T){
	dns := os.Getenv("DATABASE_URL")
	if dns == "" {
		t.Skip("DATABASE_URL is not set")
	}

	conn, err := pgx.Connect(context.Background(),dns)
	if err != nil {
		t.Fatalf("Unable to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

}