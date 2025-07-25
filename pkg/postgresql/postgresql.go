package postgresql

import (
	"github.com/jackc/pgx/v5"
	"context"
)

func ConnectToDB(url string) (*pgx.Conn, error){
	conn, err := pgx.Connect(context.Background(), url)
	return conn, err
}

func CreateDB(conn *pgx.Conn, dbName string) error {
	_, erro := conn.Exec(context.Background(), "CREATE DATABASE "+dbName)
	return erro
}