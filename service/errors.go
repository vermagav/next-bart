package service

import (
	"errors"
)

var (
	errSerializingJson    = errors.New("Error serializing JSON response")
	errParsingQueryParams = errors.New("Error parsing supplied query parameters")
)
