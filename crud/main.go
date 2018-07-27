package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"

	crud "experiments/crud/crud"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
	)
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	dbConn, err := sqlx.Connect(`postgres`, `user=vikipass dbname=vikipass sslmode=disable`)
	if err != nil {
		logger.Log("Could not open db" + err.Error())
	}

	fmt.Println("main.go: Creating service..")
	var svc = crud.NewService(dbConn)

	fmt.Println("main.go: Calling handler")
	var h http.Handler
	{
		h = crud.MakeHTTPHandler(svc, log.With(logger, "component", "HTTP"))
	}
	//crud.MakeHTTPHandler(svc, log.With(logger, "component", "HTTP"))
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, h)
	}()

	logger.Log("exit", <-errs)

	fmt.Println("Out!")
}

/*
{
  "driver": "postgres",
  "host": "192.168.99.100",
  "port": 5432,
  "enableSSL": false,
  "username": "myuser",
  "password": "",
  "databaseName": "mydb",
  "maxIdleConnections": 4,
  "maxOpenConnections": 8
}
*/
