package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (ctr *Controller) Get(w http.ResponseWriter, r *http.Request) {
	trackNumber := mux.Vars(r)["trackNumber"]

	err := validateTrackNumber(trackNumber)
	if err != nil {
		ctr.handleError(w, err)

		return
	}

	ctx := r.Context()

	order, err := ctr.apiService.Get(ctx, trackNumber)
	if err != nil {
		ctr.handleError(w, responseError{
			Kind:   "business",
			status: http.StatusInternalServerError,
			Detail: err.Error(),
		})

		return
	}

	response := toGetOrderResponse(order)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		ctr.handleError(w, responseError{
			Kind:   "business",
			status: http.StatusInternalServerError,
			Detail: err.Error(),
		})

		return
	}
}
