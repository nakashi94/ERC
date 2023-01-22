package database

import (
	// "context"

	"context"
	"fmt"
	"log"
	"web-service-api/config"

	"github.com/jackc/pgx/v5"
)

type Conn struct {
	DB *pgx.Conn
}

func NewConnection() (*Conn, error) {
	dbURL, err := config.GetDBURL()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := pgx.Connect(context.Background(), dbURL)
	fmt.Println("Connect Successfully")

	return &Conn{
		DB: conn,
	}, nil
}

func (c *Conn) Close() error {
	if err := c.Close(); err != nil {
		return err
	}
	return nil
}
