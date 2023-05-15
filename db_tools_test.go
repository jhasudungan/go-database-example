package main

import (
	"fmt"
	"log"
	"testing"
)

func TestConnect(t *testing.T) {

	dbtools := &DbToolsImpl{}

	con, err := dbtools.CreateConnection()
	defer con.Close()

	if err != nil {
		log.Fatal("Error connection to DB : " + err.Error())
	}

	con.Ping()

}

func TestCreateConnectionString(t *testing.T) {

	dbtools := &DbToolsImpl{}

	connectionString, err := dbtools.CreateConnectionString()

	if err != nil {
		log.Fatal("Error connection to DB : " + err.Error())
	}

	fmt.Println(connectionString)
}
