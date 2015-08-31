package bart

import (
	"errors"
)

var (
	errCreatingRequest = errors.New("Error creating an HTTP request to the BART API")
	errReadingBody     = errors.New("Error reading response body from request made to the BART API")
	errParsingBody     = errors.New("Error when parsing response body from request made to the BART API")
	errStationNotFound = errors.New("Error finding a station by the specified id")
)
