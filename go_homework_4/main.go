package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32 //BAD, but this is h/w

func (d dollars) String() string { //prints dollars in formatted string
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

// add handlers
func (db database) list(w http.ResponseWriter, req *http.Request) { //method on db allows us to use it as a handler
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) add(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")   //get item from url request message
	price := req.URL.Query().Get("price") //get price from url request message

	if _, ok := db[item]; ok { //get boolean whether value exists
		msg := fmt.Sprintf("duplicate item: %q", item) //provide msg string
		http.Error(w, msg, http.StatusBadRequest)      //give 400 http error
		return
	}

	p, err := strconv.ParseFloat(price, 32)

	if err != nil {
		msg := fmt.Sprintf("invalid price: %q", price)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	db[item] = dollars(p)

	fmt.Fprintf(w, "added %s with price %s\n", item, db[item])

}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")   //get item from url request message
	price := req.URL.Query().Get("price") //get price from url request message

	if _, ok := db[item]; !ok { //get boolean whether value exists
		msg := fmt.Sprintf("no such item : %q", item) //provide msg string
		http.Error(w, msg, http.StatusNotFound)       //give 400 http error
		return
	}

	p, err := strconv.ParseFloat(price, 32)

	if err != nil {
		msg := fmt.Sprintf("invalid price: %q", price)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	db[item] = dollars(p)

	fmt.Fprintf(w, "new price %s for price %s\n", db[item], item)

}

func (db database) fetch(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("no such item: %q", item)
		http.Error(w, msg, http.StatusNotFound) //404
		return
	}

	fmt.Fprintf(w, "You fetched %s item with %s price", item, db[item])

}

func (db database) drop(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if _, ok := db[item]; !ok { //check if item exists
		msg := fmt.Sprintf("no such item to delete: %q", item)
		http.Error(w, msg, http.StatusNotFound) //404
		return
	}

	delete(db, item) //deletes item from map
	fmt.Fprintf(w, "You deleted item %s", item)

}

func main() {
	db := database{
		"shoes": 50,
		"socks": 5,
	}

	//add some routes

	http.HandleFunc("/list", db.list)     //execute method at localhost:8080/list
	http.HandleFunc("/create", db.add)    //add item to database map
	http.HandleFunc("/update", db.update) //update a price that is in the database map
	http.HandleFunc("/read", db.fetch)    //retrieve an item from the database map
	http.HandleFunc("/delete", db.drop)   //delete an item from the database map

	log.Fatal(http.ListenAndServe(":8080", nil)) //create server that listens on port 8080
}
