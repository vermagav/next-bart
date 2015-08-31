package service

import (
	"strconv"
)

// parseCount takes a raw query parameter string and parses it into an int.
// If the query parameter isn't present, default to 0.
// If it is malformed, return a bad request error.
func parseCount(in string) (int, error) {
	count := 0
	var err error
	if in != "" {
		count, err = strconv.Atoi(in)
		if err != nil {
			return 0, errParsingQueryParams
		}
		if count <= 0 {
			return 0, errCountInvalid
		}
	}
	return count, nil
}

// parseLatLong takes raw query parameter strings and parses them into float64.
// If both lat and long aren't present, default to 0.
// If only is present, or they're malformed, return a bad request error.
func parseLatLong(inLat string, inLong string) (float64, float64, error) {
	lat := 0.
	long := 0.
	var err error
	if inLat != "" || inLong != "" {
		lat, err = strconv.ParseFloat(inLat, 64)
		if err != nil {
			return 0., 0., errParsingQueryParams
		}
		long, err = strconv.ParseFloat(inLong, 64)
		if err != nil {
			return 0., 0., errParsingQueryParams
		}
	}
	return lat, long, nil
}
