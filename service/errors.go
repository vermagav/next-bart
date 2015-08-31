package service

import (
	"errors"
)

var (
	errCountInvalid       = errors.New("Error parsing count query paramater; count must be greater than 0")
	errParsingQueryParams = errors.New("Error parsing supplied query parameters")
	errSerializingJson    = errors.New("Error serializing JSON response")
)
