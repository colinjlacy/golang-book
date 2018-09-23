package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

type database map[string]dollars

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	priceString := r.URL.Query().Get("price")

	price, err := strconv.Atoi(priceString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "you must include a number as a price: %q\n", price)
		return
	}
	if price == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "price must be higher than 0: %q\n", price)
		return
	}

	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	db[item] = dollars(price)

	fmt.Fprintf(w, "%s: $%.2f\n", item, db[item])
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
