package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var orders []Order

func postOrder(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var order Order
	_ = json.NewDecoder(r.Body).Decode(&order)
	orders = append(orders, order)
	json.NewEncoder(w).Encode(&order)

	time.Sleep(time.Second * 3)

	fmt.Print("\nDining-hall recieved the order:\n", order)
}

func main() {
	router := mux.NewRouter()

	//URL path and the function to handle
	router.HandleFunc("/distribution", postOrder).Methods("POST")

	for i := 0; i < 10; i++ {
		waitToOrder()
	}

	http.ListenAndServe(":3030", router)
}
