package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"

	"net/http"
)

type Item struct {
	UID string `json: "UID"`
	Name string `json: "Name"`
	Desc string `json: "Desc"`
	Price float64 `json: "Price"`
}

type User struct {
	ID uint64 `json:"id"`
	FirstName string `json:"firstName" validate:"required"`
	LastName string `json:"lastName" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	RepeatPassword string `json:"password_repeat" validate:"required"`
}

var inventory []Item
var user []User

func homePage(w http.ResponseWriter, r *http.Request ) {
	fmt.Fprintf(w, "Endpoint called: homePage()")
}

func getInventory(w http.ResponseWriter, r *http.Request){
	//w.Header().Set("Content-Type", "application/json")
	fmt.Println("Function called: getInventory")

	//json.NewEncoder(w).Encode(inventory)
}

func handleRequests(){
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/inventory", getInventory)

	router.HandleFunc("/", homePage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	inventory = append(inventory, Item{
		UID: "0",
		Name: "Cheese",
		Desc: "A fine ...",
		Price: 5,
	})
	handleRequests()


}