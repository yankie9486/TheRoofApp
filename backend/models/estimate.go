package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	Config "github.com/yankie9486/TheRoofApp/backend/config"
)

//Estimate for job
type Estimate struct {
	ID          int64     `db:"id" json:"id"`
	JobAddress  string    `db:"job_address" json:"job_address"`
	JobAddress2 string    `db:"job_address2" json:"job_address2"`
	JobCity     string    `db:"job_city" json:"job_city"`
	JobState    string    `db:"job_state" json:"job_state"`
	JobZip      string    `db:"job_zip" json:"job_zip"`
	CreateTime  time.Time `db:"create_time" json:"create_time"`
	UpdateTime  time.Time `db:"update_time" json:"update_time"`
	Status      int8      `db:"status" json:"status"`
}

//Returns estimates table name
func EstimatesTableName() string {
	return "estimates"
}

//GetAllEstimates
func GetAllEstimates(estimate *[]Estimate) (err error) {
	if err = Config.DB.Find(estimate).Error; err != nil {
		return err
	}
	return nil
}

//CreateEstimates
func CreateEstimate(estimate *Estimate) (err error) {
	if err = Config.DB.Create(estimate).Error; err != nil {
		return err
	}
	return nil
}

//GetEstimate
func GetEstimate(estimate *Estimate, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(estimate).Error; err != nil {
		return err
	}
	return nil
}

//UpdateEstimate
func UpdateEstimate(estimate *Estimate, id string) (err error) {
	if err = Config.DB.Save(estimate).Error; err != nil {
		return err
	}
	return nil
}

//DeleteEstimate
func DeleteEstimate(estimate *Estimate, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).Delete(estimate).Error; err != nil {
		return err
	}
	return nil
}
