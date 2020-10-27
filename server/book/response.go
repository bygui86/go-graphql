package book

import (
	"encoding/json"
	"net/http"
	"time"
)

// TODO maybe for REST APIs?

func HttpResponseSuccess(w http.ResponseWriter, r *http.Request, data interface{}) {
	setResponse := SetResponse{
		Status:     http.StatusText(200),
		AccessTime: time.Now().Format("02-01-2006 15:04:05"),
		Data:       data,
	}
	response, jsonErr := json.Marshal(setResponse)
	if jsonErr != nil {
		// TODO
	}

	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(200)

	_, wErr := w.Write(response)
	if wErr != nil {
		// TODO
	}
}

func HttpResponseError(w http.ResponseWriter, r *http.Request, data interface{}, code int) {
	setResponse := SetResponse{
		Status:     http.StatusText(code),
		AccessTime: time.Now().Format("02-01-2006 15:04:05"),
		Data:       data,
	}
	response, jsonErr := json.Marshal(setResponse)
	if jsonErr != nil {
		// TODO
	}

	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(code)

	_, wErr := w.Write(response)
	if wErr != nil {
		// TODO
	}
}
