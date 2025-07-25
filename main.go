package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ninoamine/A-S-Controllers---Postgresql/pkg/postgresql"
)


func main() {
	url := os.Getenv("DATABASE_URL")
	conn, err := postgresql.ConnectToDB(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected to database at %s\n", url)
	fmt.Println("Creating database 'testdb'...")
	if err := postgresql.CreateDB(conn,`"testdb-controller"`); err != nil {
		panic(err)
	}
	fmt.Println("Database 'testdb' created successfully.")
	defer conn.Close(context.Background())
}