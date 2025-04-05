package controller

import (
	"encoding/json"
	"net/http"
)

func (ctr *Controller) Get(w http.ResponseWriter, r *http.Request) {
	trackNumber := r.URL.Query().Get("track_number")

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

	err = json.NewEncoder(w).Encode(order)
	if err != nil {
		ctr.handleError(w, responseError{
			Kind:   "business",
			status: http.StatusInternalServerError,
			Detail: err.Error(),
		})

		return
	}
}
