package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yankie9486/TheRoofApp/backend/models"
)

//GetEstimates
func GetAllEstimates(c *gin.Context) {
	var estimate []models.Estimate

	err := models.GetAllEstimates(&estimate)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, estimate)
	}
}

//CreateEstimate
func CreateEstimate(c *gin.Context) {
	var estimate models.Estimate

	c.BindJSON(&estimate)
	fmt.Println(estimate)
	err := models.CreateEstimate(&estimate)
	if err != nil {
		// c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, estimate)
	}
}

//GetEstimate
func GetEstimate(c *gin.Context) {
	id := c.Params.ByName("id")
	var estimate models.Estimate

	err := models.GetEstimate(&estimate, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, estimate)
	}
}

//UpdateEstimate
func UpdateEstimate(c *gin.Context) {
	id := c.Params.ByName("id")
	var estimate models.Estimate

	err := models.GetEstimate(&estimate, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.BindJSON(&estimate)
	err = models.UpdateEstimate(&estimate, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, estimate)
	}
}

//GetEstimate
func DeleteEstimate(c *gin.Context) {
	id := c.Params.ByName("id")
	var estimate models.Estimate

	err := models.DeleteEstimate(&estimate, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
