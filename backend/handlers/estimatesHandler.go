package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/yankie9486/TheRoofApp/backend/estimate"
	"gopkg.in/mgo.v2/bson"
)

func bodyToEstimate(r *http.Request, e *estimate.Estimate) error {
	if r.Body == nil {
		return errors.New("Request Body empty")
	}

	if e == nil {
		return errors.New("A Estimate is required")
	}

	bd, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(bd, e)

}

func estimatesGetAll(w http.ResponseWriter, r *http.Request) {
	estimates, err := estimate.All()
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}

	postBodyResponse(w, http.StatusOK, jsonRepsonse{"estimates": estimates})

}

func estimatePostOne(w http.ResponseWriter, r *http.Request) {
	e := new(estimate.Estimate)
	err := bodyToEstimate(r, e)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	e.ID = bson.NewObjectId()
	err = e.Save()
	if err != nil {
		if err == estimate.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Location", "/users/"+e.ID.Hex())
	w.WriteHeader(http.StatusCreated)
}
