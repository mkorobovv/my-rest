package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (ctr *Controller) handleError(w http.ResponseWriter, err error) {
	respErr := new(responseError)

	switch {
	case errors.As(err, respErr):
		w.WriteHeader(respErr.status)

		respJsonBytes, marshalErr := json.Marshal(respErr)
		if marshalErr != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		_, writeErr := w.Write(respJsonBytes)
		if writeErr != nil {
			ctr.logger.Error(writeErr.Error())
		}
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	return
}

const trackNumberLength = 20

func validateTrackNumber(trackNumber string) error {
	if len(trackNumber) != trackNumberLength {
		respErr := responseError{
			Kind:   "validation",
			status: http.StatusBadRequest,
			Detail: fmt.Sprintf("invalid length for track_number: [%s]", trackNumber),
		}

		return respErr
	}

	return nil
}
