package controller

import (
	"encoding/json"
	"net/http"
)

func (ctr *Controller) Update(w http.ResponseWriter, r *http.Request) {
	trackNumber := r.URL.Query().Get("track_number")

	err := validateTrackNumber(trackNumber)
	if err != nil {
		ctr.handleError(w, err)

		return
	}

	var dtoIn OrderDTO

	err = json.NewDecoder(r.Body).Decode(&dtoIn)
	if err != nil {
		ctr.handleError(w, responseError{
			Kind:   "validation",
			status: http.StatusBadRequest,
			Detail: err.Error(),
		})

		return
	}

	err = dtoIn.validate()
	if err != nil {
		ctr.handleError(w, err)

		return
	}

	ctx := r.Context()

	request := dtoIn.toRequest()

	err = ctr.apiService.Update(ctx, request)
	if err != nil {
		ctr.handleError(w, responseError{
			Kind:   "business",
			status: http.StatusInternalServerError,
			Detail: err.Error(),
		})

		return
	}
}
