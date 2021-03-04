package estimate

import (
	"errors"

	storm "github.com/asdine/storm/v3"
	"gopkg.in/mgo.v2/bson"
)

//Estimate for job
type Estimate struct {
	ID         bson.ObjectId `json:"id"`
	JobAddress Address       `json:"jobaddress"`
	RoofType   []RoofType    `json:"rooftypes"`
}

//Address type
type Address struct {
	ID       bson.ObjectId `json:"id"`
	Address  string        `json:"address"`
	Address2 string        `json:"address2"`
	City     string        `json:"city"`
	State    string        `json:"state"`
	Zip      string        `json:"zip"`
}

//Area of the square measurement
type Area struct {
	ID     bson.ObjectId `json:"id"`
	Width  float64       `json:"width"`
	Height float64       `json:"height"`
	Total  float64       `json:"total"`
}

//RoofCovering is type of roof top product ex. Shingle, Tile, Flat
type RoofCovering struct {
	ID   bson.ObjectId
	Name string
}

//RoofType combine the Area and roof covering
type RoofType struct {
	ID           bson.ObjectId `json:"id"`
	AreaList     []Area        `json:"areas"`
	RoofCovering RoofCovering  `json:"roofCovering"`
}

const (
	dbPath = "estimate.db"
)

// errors
var (
	ErrRecordInvalid = errors.New("record is invaild")
)

//All get all estimates
func All() ([]Estimate, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	estimates := []Estimate{}
	err = db.All(&estimates)
	if err != nil {
		return nil, err
	}
	return estimates, nil
}

//One get estimate by id
func One(id bson.ObjectId) (*Estimate, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	e := new(Estimate)
	err = db.One("ID", id, e)
	if err != nil {
		return nil, err
	}
	return e, nil
}

//Delete estimate by id
func Delete(id bson.ObjectId) error {
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	e := new(Estimate)
	err = db.One("ID", id, e)
	if err != nil {
		return err
	}

	return db.DeleteStruct(e)
}

//Save updates or creates a given record in the database
func (e *Estimate) Save() error {
	if err := e.validate(); err != nil {
		return err
	}
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Save(e)
}

//validate make sure that the record contains valid data
func (e *Estimate) validate() error {
	if e.JobAddress.Address == "" {
		return ErrRecordInvalid
	}
	return nil
}
