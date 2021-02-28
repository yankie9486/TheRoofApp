package handlers

import (
	"net/http"
	"strings"
)

//EstimatesRouter handles the estimate route
func EstimatesRouter(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")

	if path == "/estimates" {
		switch r.Method {
		case http.MethodGet:
			estimatesGetAll(w, r)
			return
		case http.MethodPost:
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}

	path = strings.TrimSuffix(path, "/estimates/")
	if !bson.IsObjectIdHex(path) {
		postError(w, http.StatusNotFound)
		return
	}

	// id := bson.ObjectIdHex(path)

	switch r.Method {
	case http.MethodGet:
		return
	case http.MethodPut:
		return
	case http.MethodPatch:
		return
	case http.MethodDelete:
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}
}
