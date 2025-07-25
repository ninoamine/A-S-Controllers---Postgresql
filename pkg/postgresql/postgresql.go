package postgresql

import (
	"github.com/jackc/pgx/v5"
	"context"
)

func ConnectToDB (url string) (*pgx.Conn, error){
	conn, err := pgx.Connect(context.Background(), url)
	return conn, err
}