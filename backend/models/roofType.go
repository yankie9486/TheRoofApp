package models

//RoofType combine the Area and roof covering
type RoofType struct {
	ID               int64           `json:"id"`
	ExistingCovering RoofCovering    `json:"existingCovering"`
	RoofMeasurement  RoofMeasurement `json:"roofMeasurement"`
	SlopeType        string
}

//RoofMeasurement Get roof measurement
type RoofMeasurement struct {
	AreaList      []Area  `json:"areas"`
	RidgeTotal    float32 `json:"ridgeTotal"`
	GableEndTotal float32 `json:"gableEndTotal"`
	DripEdgeTotal float32 `json:"dripEdgeTotal"`
	Slope         float32 `json:"slope"`
	ValleyTotal   float32 `json:"valleyTotal"`
	Height        float32 `json:"height"`
}

//Area of the square measurement
type Area struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Total  float64 `json:"total"`
}

//RoofCovering is type of roof top product ex. Shingle, Tile, Flat
type RoofCovering struct {
	ID   int64
	Name string
}

const (
	//Low Slope const
	LOWSlOPE = "LOW"
	//Steep Slope const
	STEEPSLOPE = "STEEP"
)
