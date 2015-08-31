package service

import (
	"net/http"
	"strconv"

	"github.com/vermagav/next-bart/bart"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func init() {
	http.Handle("/", goji.DefaultMux)

	// Set up web handlers
	goji.Get("/", handler)
	goji.Get("/stations", stations)
	goji.Get("/stations/:stationId", stationsId)
}

func handler(c web.C, w http.ResponseWriter, r *http.Request) {
	stations, err := bart.GetStations(r, 0)
	if err != nil {
		BuildResponseInternalServerError(w, r, err)
		return
	} else {
		buildResponseSuccess(w, r, stations)
		return
	}
}

func stations(c web.C, w http.ResponseWriter, r *http.Request) {
	// Grab query parameters
	err := r.ParseForm()
	if err != nil {
		BuildResponseBadRequest(w, r, errParsingQueryParams)
		return
	}
	count, err := strconv.Atoi(r.Form.Get("count"))
	if err != nil {
		BuildResponseBadRequest(w, r, errParsingQueryParams)
		return
	}
	// lat := r.Form.Get("lat")
	// long := r.Form.Get("long")

	// Fetch station data and return
	stations, err := bart.GetStations(r, count)
	if err != nil {
		BuildResponseInternalServerError(w, r, err)
		return
	} else {
		buildResponseSuccess(w, r, stations)
		return
	}
}

func stationsId(c web.C, w http.ResponseWriter, r *http.Request) {
	id := c.URLParams["stationId"]
	station, err := bart.GetStationbyId(r, id)
	if err != nil {
		BuildResponseBadRequest(w, r, err)
		return
	} else {
		buildResponseSuccess(w, r, station)
		return
	}
}
