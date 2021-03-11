package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/asdine/storm/v3"
	"github.com/yankie9486/TheRoofApp/backend/estimate"
	"gopkg.in/mgo.v2/bson"
)

func bodyToEstimate(r *http.Request, e *estimate.Estimate) error {
	if r == nil {
		return errors.New("a request is required")
	}
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

	if r.Method == http.MethodHead {
		postBodyResponse(w, http.StatusOK, jsonRepsonse{"estimates": estimates})
	}

	postBodyResponse(w, http.StatusOK, jsonRepsonse{})

}

func estimateGetOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	e, err := estimate.One(id)
	if err != nil {
		if err == estimate.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	if r.Method == http.MethodHead {
		postBodyResponse(w, http.StatusOK, jsonRepsonse{"estimates": e})
	}
	postBodyResponse(w, http.StatusOK, jsonRepsonse{"estimate": e})
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

	w.Header().Set("Location", "/estimate/"+e.ID.Hex())
	w.WriteHeader(http.StatusCreated)
}

//You have to put all fields for it to work
func estimatePutOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	e := new(estimate.Estimate)
	err := bodyToEstimate(r, e)

	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	e.ID = id
	err = e.Save()
	if err != nil {
		if err == estimate.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	postBodyResponse(w, http.StatusOK, jsonRepsonse{"estimates": e})
}

//You have to put all fields for it to work
func estimatePatchOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	e, err := estimate.One(id)
	if err != nil {
		if err == estimate.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}

	err = bodyToEstimate(r, e)

	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	e.ID = id
	err = e.Save()
	if err != nil {
		if err == estimate.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	postBodyResponse(w, http.StatusOK, jsonRepsonse{"estimates": e})

}

func estimateDeleteOne(w http.ResponseWriter, _ *http.Request, id bson.ObjectId) {
	err := estimate.Delete(id)
	if err != nil {
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
