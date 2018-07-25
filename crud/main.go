package main

import (
	"fmt"
	"log"

	crud "experiments/crud/crud"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	dbConn, err := sqlx.Connect(`postgres`, `user=myuser dbname=mydb sslmode=disable`)
	if err != nil {
		log.Fatalf("Could not open db" + err.Error())
	}

	fmt.Println("Creating service..")
	var svc = crud.NewService(dbConn)

	fmt.Println("Calling handler")
	crud.MakeHandler(svc)

	fmt.Println("Out!")
}

/*
{
  "driver": "postgres",
  "host": "localhost",
  "port": 5432,
  "enableSSL": false,
  "username": "myuser",
  "password": "",
  "databaseName": "mydb",
  "maxIdleConnections": 4,
  "maxOpenConnections": 8
}
*/
