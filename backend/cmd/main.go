package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/yankie9486/TheRoofApp/backend/config"
	"github.com/yankie9486/TheRoofApp/backend/models"
	Routes "github.com/yankie9486/TheRoofApp/backend/routes"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql",
		config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Estimate{})

	r := Routes.SetupRouter()
	// running
	r.Run(":5000")

}
