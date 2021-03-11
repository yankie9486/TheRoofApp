package handlers

import (
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"
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
			estimatePostOne(w, r)
			return
		case http.MethodHead:
			estimatesGetAll(w, r)
			return
		case http.MethodOptions:
			postOptionsResponse(w, []string{http.MethodGet, http.MethodPost, http.MethodOptions}, nil)
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}

	path = strings.TrimPrefix(path, "/estimates/")

	if !bson.IsObjectIdHex(path) {
		postError(w, http.StatusNotFound)
		return
	}

	id := bson.ObjectIdHex(path)

	switch r.Method {
	case http.MethodGet:
		estimateGetOne(w, r, id)
		return
	case http.MethodPut:
		estimatePutOne(w, r, id)
		return
	case http.MethodPatch:
		estimatePatchOne(w, r, id)
		return
	case http.MethodDelete:
		estimateDeleteOne(w, r, id)
		return
	case http.MethodHead:
		estimateGetOne(w, r, id)
		return
	case http.MethodOptions:
		postOptionsResponse(w, []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions}, nil)
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}
}
