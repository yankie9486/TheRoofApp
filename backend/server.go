package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

// Customer information
type Customer struct {
	ID         int64   `json:"customerId"`
	FirstName  string  `json:"fname"`
	LastName   string  `json:"lname"`
	MiddleName string  `json:"mname"`
	Company    string  `json:"company"`
	Contact    Contact `json:"contact"`
	Address    Address `json:"address"`
}

//Contact info type
type Contact struct {
	ID    int64  `json:"id"`
	Phone string `json:"phone"`
	Cell  string `json:"cell"`
	Email string `json:"email"`
}

//Address type
type Address struct {
	ID       int64  `json:"id"`
	Address  string `json:"address"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zip      string `json:"zip"`
}

//Area of the square measurement
type Area struct {
	ID     int64   `json:"id"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Total  float64 `json:"total"`
}

//RoofCovering is type of roof top product ex. Shingle, Tile, Flat
type RoofCovering struct {
	ID   int64
	Name string
}

//RoofType combine the Area and roof covering
type RoofType struct {
	ID           int64        `json:"id"`
	AreaList     []Area       `json:"areas"`
	RoofCovering RoofCovering `json:"roofCovering"`
}

//Product struct for estimates
type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Cost        float64 `json:"cost"`
	Price       float64 `json:"price"`
	Tax         Tax     `json:"tax"`
}

//Tax rates for products
type Tax struct {
	ID    int64   `json:"id"`
	State string  `json:"state"`
	City  string  `json:"city"`
	Rate  float64 `json:"rate"`
}

//Estimate for job
type Estimate struct {
	ID         int64      `json:"id"`
	Customer   Customer   `json:"customers"`
	JobAddress Address    `json:"jobaddress"`
	RoofType   []RoofType `json:"rooftypes"`
}

func getContacts(w http.ResponseWriter, r *http.Request) {

	// keys, ok :=

	// decoder := json.NewDecoder(r.Body)

	var customerList []Customer

	customerList1 := Customer{
		ID:         1,
		FirstName:  "Giancarlo",
		MiddleName: "",
		LastName:   "Cunis",
		Company:    "Cunis Inc",
	}

	customerList = append(customerList, customerList1)

	customerList2 := Customer{
		ID:         2,
		FirstName:  "Giana",
		MiddleName: "",
		LastName:   "Cunis",
		Company:    "Cunis Inc",
	}

	customerList = append(customerList, customerList2)

	customerList3 := Customer{
		ID:         3,
		FirstName:  "GianLucca",
		MiddleName: "",
		LastName:   "Cunis",
		Company:    "Cunis Inc",
	}
	customerList = append(customerList, customerList3)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(customerList); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func getContactByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ID: %v\n", vars["id"])
}

func createContact(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var newCustomer Customer

	decoder.Decode(&newCustomer)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(newCustomer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func estimate(w http.ResponseWriter, request *http.Request) {

	urlParms := mux.Vars(request)
	estimateParms := urlParms["estimate"]

	// customer := Customer{}
	// jobAddress := Address{}
	// roofType := []RoofType{}

	// estimate := Estimate{
	// 	Customer:   customer,
	// 	JobAddress: jobAddress,
	// 	RoofType:   roofType,
	// }

	output, err := json.Marshal(estimateParms)
	if err != nil {
		fmt.Println("Something went wrong!")
	}

	fmt.Fprintf(w, string(output))

	// decoder := json.NewDecoder(request.Body)

	// var estimate Estimate

	// decoder.Decode(&estimate)

	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.WriteHeader(http.StatusOK)
	// if err := json.NewEncoder(w).Encode(estimate); err != nil {
	// 	panic(err)
	// }
}

func createEstimate(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var newEstimate Estimate

	decoder.Decode(&newEstimate)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(newEstimate); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	routes := mux.NewRouter()
	routes.HandleFunc("/api/estimates", estimate).Methods("GET")
	routes.HandleFunc("/api/estimates", createEstimate).Methods("POST")
	routes.HandleFunc("/api/contacts", getContacts).Methods("GET")
	routes.HandleFunc("/api/contacts/{id:[0-9]+}", getContactByID).Methods("GET")
	routes.HandleFunc("/api/contacts", createContact).Methods("POST")
	http.Handle("/", routes)

	err := http.ListenAndServe(":8020", routes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
