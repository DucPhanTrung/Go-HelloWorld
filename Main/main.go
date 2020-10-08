package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	initServer()
}

func initServer() {
	http.HandleFunc("/summary", handle)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	numbers, ok := r.URL.Query()["number"]

	if !ok || len(numbers[0]) < 1 {
		log.Println("Url Param 'number' is missing")
		return
	}

	number, err := strconv.Atoi(numbers[0])
	if err != nil {
		fmt.Println("Invalid number: ", err)
		return
	}

	numberRepository := NumberRepository{number}
	sum, err := numberRepository.SumAllPrimesNumber()
	fmt.Fprintf(w, fmt.Sprintf("Sum all primes number: %d", sum))
}
