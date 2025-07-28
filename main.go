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
	fmt.Println("Creating database 'testdb-controller'...")
	if err := postgresql.CreateDB(conn,`"testdb-controller"`); err != nil {
		panic(err)
	}
	fmt.Println("Database 'testdb-controller' created successfully.")

	if err := postgresql.DeleteDB(conn, `"testdb-controller"`); err != nil {
		panic(err)
	}
	fmt.Println("Database 'testdb-controller' deleted successfully.")
	defer conn.Close(context.Background())
}