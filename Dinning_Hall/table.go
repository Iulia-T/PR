package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Table struct {
	Id    int
	State string
	Order Order
}

var tables []Table

const (
	free           = "free"
	waitingToOrder = "waitingToOrder"
	nrOfTables     = 10
)

func genTable(id int) Table {
	return Table{
		Id:    id,
		State: free,
		Order: genOrder(),
	}
}

func initTables() {
	for i := 0; i < nrOfTables; i++ {
		newTable := genTable(i + 1)
		tables = append(tables, newTable)
	}
}

func performPostRequest() {
	const myUrl = "http://localhost:8080/order"

	for _, table := range tables {
		if table.State == free {
			table.State = waitingToOrder
			fmt.Println("Table", table.Id, "is ready to make an order.")
		}

		time.Sleep(time.Second * 3)

		requestBody, _ := json.Marshal(table.Order)

		response, err := http.Post(myUrl, "application/json", bytes.NewBuffer(requestBody))

		if err != nil {
			panic(err)
		}

		defer response.Body.Close()

		fmt.Printf("\nOrder sent to kitchen\n")
	}
}

func waitToOrder() {
	go performPostRequest()
}
