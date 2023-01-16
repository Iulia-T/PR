package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func postOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var order Order

	_ = json.NewDecoder(r.Body).Decode(&order)

	// add orders to the end of the queue
	orderList.Enqueue(order)

	json.NewEncoder(w).Encode(&order)

	fmt.Print("\nKitchen recieved order:\n", order)

	go performPostRequest(order)
}

func performPostRequest(order Order) {
	if !orderList.isEmpty() {
		const myUrl = "http://localhost:9000/distribution"

		// return the first order form the queue
		order := orderList.Dequeue()

		var requestBody, _ = json.Marshal(order)

		fmt.Println("\nOrder with id ", order.Id, " is being cooked.")
		time.Sleep(time.Second * 3)

		fmt.Printf("\nOrder %v was sent to New \n", string(requestBody))
		response, err := http.Post(myUrl, "application/json", bytes.NewBuffer(requestBody))

		if err != nil {
			panic(err)
		}
		defer response.Body.Close()

		time.Sleep(time.Second * 1)
	}
}

func main() {

	router := mux.NewRouter()

	//URL path and the function to handle
	router.HandleFunc("/order", postOrder).Methods("POST")

	http.ListenAndServe(":8080", router)
}
