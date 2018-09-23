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
	switch request.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(writer, "%s: %s\n", item, price)
		}
	case "/price":
		item := request.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			writer.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(writer, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(writer, "%s\n", price)
	default:
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(writer, "no such page: %s\n", request.URL)
		return
	}
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
