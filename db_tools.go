package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DbTools interface {
	CreateConnectionString() (string, error)
	CreateConnection() (*sql.DB, error)
	CheckConnectionString() error
}

type DbToolsImpl struct{}

func (d *DbToolsImpl) CreateConnectionString() (string, error) {
	err := godotenv.Load()

	if err != nil {
		return "", nil
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_DATABASE_NAME")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	connectionString := fmt.Sprintf("host=%v port=%v user=%v "+
		"password=%v dbname=%v sslmode=disable",
		host, port, username, password, dbname)

	return connectionString, nil
}

func (d *DbToolsImpl) CreateConnection() (*sql.DB, error) {

	connectionString, err := d.CreateConnectionString()

	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *DbToolsImpl) CheckConnectionString() error {

	connectionString, err := d.CreateConnectionString()

	if err != nil {
		return err
	}

	log.Printf("Connection Status : %v ", connectionString)

	return err
}
