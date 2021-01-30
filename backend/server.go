package main

import (
	"encoding/json"
	"net/http"
)

// Customer information
type Customer struct {
	FirstName  string
	LastName   string
	MiddleName string
	Company    string
	Contact    []Contact
	Address    []Address
}

//Contact info type
type Contact struct {
	Phone string
	Cell  string
	Email string
}

//Address type
type Address struct {
	Address  string
	Address2 string
	City     string
	State    string
	Zip      string
}

//Area of the square measurement
type Area struct {
	Width  float64
	Height float64
	Total  float64
}

//RoofType combine the Area and roof covering
type RoofType struct {
	AreaList     []Area
	RoofCovering []string
}

//Estimate for job
type Estimate struct {
	Customer   Customer
	JobAddress Address
	RoofType   []RoofType
}

func estimate(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var estimate Estimate

	decoder.Decode(&estimate)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(estimate); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/estimate", estimate)
	http.ListenAndServe(":8090", nil)
}
