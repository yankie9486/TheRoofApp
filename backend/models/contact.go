package models

type Person struct {
	FirstName  string  `json:"firstName"`
	MiddleName string  `json:"middleName"`
	LastName   string  `json:"lastName"`
	Company    string  `json:"companyName"`
	WorkPhone  string  `json:"workPhone"`
	HomePhone  string  `json:"homePhone"`
	CellPhone  string  `json:"cellPhone"`
	Email      string  `json:"email"`
	AddressID  float64 `json:"addressID"`
}

type Contact struct {
	Person []Person `json:"persons"`
	Status string   `json:"status"`
}
