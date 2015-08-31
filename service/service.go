package service

import (
	"net/http"
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

	// Parse count
	count, err := parseCount(r.Form.Get("count"))
	if err != nil {
		BuildResponseBadRequest(w, r, err)
	}

	// Parse lat/long
	lat, long, err := parseLatLong(r.Form.Get("lat"), r.Form.Get("long"))
	if err != nil {
		BuildResponseBadRequest(w, r, err)
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
