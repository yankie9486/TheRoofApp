package handlers

import (
	"net/http"

	"backend/estimate"
)

func estimatesGetAll(w http.ResponseWriter, r *http.Request) {
	estimates, err := estimate.All()
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}

	postBodyResponse(w, http.StatusOK, jsonRepsonse{"estimates": estimates})

}
