package service

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/vermagav/next-bart/bart"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func init() {
	http.Handle("/", goji.DefaultMux)

	// Set up handlers for endpoints
	goji.Get("/stations", stations)
	goji.Get("/stations/:stationId", stationsId)
	goji.Get("/stations/:stationId/departures", departures)
}

func stations(c web.C, w http.ResponseWriter, r *http.Request) {
	// Grab query parameters and parse into map
	err := r.ParseForm()
	if err != nil {
		BuildResponseBadRequest(w, r, errParsingQueryParams)
		return
	}

	// Parse count: if included, try to convert to int, default to 0
	count := 0
	countRaw := r.Form.Get("count")
	if countRaw != "" {
		// If the count param was included, try to convert to int
		count, err = strconv.Atoi(countRaw)
		if err != nil {
			BuildResponseBadRequest(w, r, errParsingQueryParams)
			return
		}
		if count <= 0 {
			BuildResponseBadRequest(w, r, errCountInvalid)
			return
		}
	}

	// Parse lat/long: if included, try to convert to float, default to 0.0f
	lat := 0.
	long := 0.
	latRaw := r.Form.Get("lat")
	longRaw := r.Form.Get("long")
	if latRaw != "" || longRaw != "" {
		lat, err = strconv.ParseFloat(latRaw, 64)
		if err != nil {
			BuildResponseBadRequest(w, r, errParsingQueryParams)
			return
		}
		long, err = strconv.ParseFloat(longRaw, 64)
		if err != nil {
			BuildResponseBadRequest(w, r, errParsingQueryParams)
			return
		}
	}

	// Fetch station data and return
	stations, err := bart.GetStations(r, count, lat, long)
	if err != nil {
		BuildResponseInternalServerError(w, r, err)
		return
	} else {
		buildResponseSuccess(w, r, stations)
		return
	}
}

func stationsId(c web.C, w http.ResponseWriter, r *http.Request) {
	// Grab URL parameter
	id := c.URLParams["stationId"]
	id = strings.ToUpper(id)

	// Fetch station data and return
	station, err := bart.GetStationbyId(r, id)
	if err != nil {
		BuildResponseBadRequest(w, r, err)
		return
	} else {
		buildResponseSuccess(w, r, station)
		return
	}
}

func departures(c web.C, w http.ResponseWriter, r *http.Request) {
	// Grab URL parameter
	id := c.URLParams["stationId"]
	id = strings.ToUpper(id)

	// Fetch departure information and return
	departures, err := bart.GetDepartures(r, id)
	if err != nil {
		BuildResponseBadRequest(w, r, err)
		return
	} else {
		buildResponseSuccess(w, r, departures)
		return
	}
}
