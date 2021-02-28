package main

import (
	"fmt"
	"net/http"
	"os"

	"./handlers"
)

func main() {

	http.HandleFunc("/estimates/", handlers.EstimatesRouter)
	http.HandleFunc("/estimates", handlers.EstimatesRouter)
	http.HandleFunc("/", handlers.RootHandler)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
