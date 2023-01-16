package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/sync/semaphore"
)

func postOrder1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var order Order

	_ = json.NewDecoder(r.Body).Decode(&order)

	// add orders to the end of the queue
	orderList.Enqueue(order)

	json.NewEncoder(w).Encode(&order)

	fmt.Print("\n New recieved order:\n", order)

}

func SendToKitchen() {
	for {
		for !orderList.isEmpty() {
			go performPostRequest1()
			time.Sleep(time.Second * 1)
		}
	}

}

func performPostRequest1() {
	const myUrl = "http://localhost:8080/distribution"

	sem := semaphore.NewWeighted(10)

	if err := sem.Acquire(context.Background(), 1); err != nil {
		log.Fatal(err)
	}

	// return the first order form the queue
	order := orderList.Dequeue()

	defer sem.Release(1)

	var requestBody, _ = json.Marshal(order)

	fmt.Println("\nOrder with id ", order.Id, " is being prepared.")
	time.Sleep(time.Second * 3)

	fmt.Printf("\nOrder %v was sent to the kitchen\n", string(requestBody))
	response, err := http.Post(myUrl, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

}

func postOrder2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var order Order

	_ = json.NewDecoder(r.Body).Decode(&order)

	// add orders to the end of the queue
	orderList2.Enqueue2(order)

	json.NewEncoder(w).Encode(&order)

	fmt.Print("\n New recieved order:\n", order)

}

func SendToDinningHall() {
	for {
		for !orderList2.isEmpty2() {
			go performPostRequest2()
			time.Sleep(time.Second * 1)
		}
	}
}

func performPostRequest2() {
	const myUrl = "http://localhost:3030/distribution"

	sem := semaphore.NewWeighted(10)

	if err := sem.Acquire(context.Background(), 1); err != nil {
		log.Fatal(err)
	}

	// return the first order form the queue
	order := orderList2.Dequeue2()

	defer sem.Release(1)

	var requestBody, _ = json.Marshal(order)

	fmt.Printf("\nOrder %v was sent to the dining-hall\n", string(requestBody))
	response, err := http.Post(myUrl, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
}

func main() {

	router := mux.NewRouter()

	//URL path and the function to handle

	router.HandleFunc("/order", postOrder1).Methods("POST")
	router.HandleFunc("/order", postOrder2).Methods("POST")

	go SendToKitchen()
	go SendToDinningHall()

	http.ListenAndServe(":9000", router)

}
