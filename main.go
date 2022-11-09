package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"

	"github.com/haagor/orderMP/adapter"
	postgresDB "github.com/haagor/orderMP/adapter"
	"github.com/haagor/orderMP/controller"
	"github.com/haagor/orderMP/model"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		postgresDB.Host, postgresDB.Port, postgresDB.User, postgresDB.Password, postgresDB.Dbname)

	var err error
	var db *sql.DB
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	dba := postgresDB.PostgresAdapter{db}

	workers := flag.Int("workers", runtime.NumCPU(), "workers")
	flag.Parse()

	ordersChan := make(chan string, 0)
	var wg sync.WaitGroup

	go func(wg *sync.WaitGroup) {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		close(ordersChan)
		wg.Wait()
		panic(1)
	}(&wg)

	for i := 0; i < *workers; i++ {
		go processOrder(ordersChan, dba, &wg)
	}

	http.Handle("/ticket", controller.OrderHandler(ordersChan, dba))
	http.ListenAndServe(":8080", nil)
}

func processOrder(ordersChan chan string, pa adapter.PostgresAdapter, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	for s := range ordersChan {
		order, err := model.StringToOrder(s)
		if err != nil {
			fmt.Printf("failed to convert order to string: %s\n", err)
			return
		}

		err = pa.AddOrderWithProduct(order)
		if err != nil {
			fmt.Printf("failed to add new order to db: %s\n", err)
			return
		}
	}
}
