package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) add(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")

	if _, ok := db[item]; ok {
		msg := fmt.Sprintf("dupicate item: %q", item)
		http.Error(w, msg, http.StatusBadRequest) //400
		return
	}
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		msg := fmt.Sprintf("invalid price: %q", price)
		http.Error(w, msg, http.StatusBadRequest) //400
		return
	}

	db[item] = dollars(p)
	fmt.Fprintf(w, "added %s with price %s\n", item, db[item])
}

func (db database) update(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("no such item: %q", item)
		http.Error(w, msg, http.StatusNotFound) //404
		return
	}
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		msg := fmt.Sprintf("invalid price: %q", price)
		http.Error(w, msg, http.StatusBadRequest) //400
		return
	}

	db[item] = dollars(p)
	fmt.Fprintf(w, "new price %s for item %s\n", db[item], item)
}

func (db database) drop(w http.ResponseWriter, r *http.Request) {
	for item := range db {
		delete(db, item)
	}
	msg := fmt.Sprintf("database is clear")
	http.Error(w, msg, http.StatusOK)
}

func (db database) fetch(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("no such item: %q", item)
		http.Error(w, msg, http.StatusNotFound) //404
		return
	}

	msg := fmt.Sprintf("price of %s: %s", item, db[item])
	http.Error(w, msg, http.StatusOK)
}

func (db database) delete(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("no such item: %q", item)
		http.Error(w, msg, http.StatusNotFound) //404
		return
	}

	price := db[item]

	delete(db, item)
	msg := fmt.Sprintf("item %s with price %s is deleted", item, price)
	http.Error(w, msg, http.StatusOK)
}

func main() {
	db := database{
		"shoes": 50,
		"socks": 5,
	}

	http.HandleFunc("/list", db.list)
	http.HandleFunc("/create", db.add)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/drop", db.drop)
	http.HandleFunc("/read", db.fetch)
	http.HandleFunc("/delete", db.delete)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
