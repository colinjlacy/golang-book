package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

type database map[string]dollars

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func (db database) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	for item, price := range db {
		fmt.Fprintf(writer, "%s: %s\n", item, price)
	}
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
