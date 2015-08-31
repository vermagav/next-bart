package service

import (
	"encoding/json"
	"io"
	"net/http"
)

// Normalize all responses to fit inside this outer schema
type Response struct {
	Data  interface{} `json:"Data"`
	Error interface{} `json:"error"`
}

// Build a success response with HTTP status 200 and json response data
func buildResponseSuccess(w http.ResponseWriter, r *http.Request, data interface{}) {
	// Set up JSON response
	body, err := json.Marshal(Response{
		Data:  data,
		Error: nil,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, errSerializingJson.Error())
		return
	}

	// Set headers
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// Write body
	io.WriteString(w, string(body))
}

// Build a failure response with HTTP status 400 and json error data
func BuildResponseBadRequest(w http.ResponseWriter, r *http.Request, err error) {
	// Set up JSON response
	body, err := json.Marshal(Response{
		Data:  nil,
		Error: err.Error(),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, errSerializingJson.Error())
		return
	}

	// Set headers
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)

	// Write body
	io.WriteString(w, string(body))
}

// Build a failure response with HTTP status 500 and json error data
func BuildResponseInternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	// Set up JSON response
	body, err := json.Marshal(Response{
		Data:  nil,
		Error: err.Error(),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, errSerializingJson.Error())
		return
	}

	// Set headers
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)

	// Write body
	io.WriteString(w, string(body))
}
