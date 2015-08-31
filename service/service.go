package service

import (
	"fmt"
	"net/http"

	"github.com/vermagav/next-bart/bart"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func init() {
	http.Handle("/", goji.DefaultMux)

	// Set up web handlers
	goji.Get("/", handler)
	goji.Get("/name/:name", handlerName)
}

func handler(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, web service!\n\n")

	stations, err := bart.GetStations(r)
	if err != nil {
		fmt.Fprint(w, "\nError getting stations: "+err.Error())
	} else {
		for k, v := range stations {
			fmt.Fprintf(w, "\n%s: %s, %v, %v", k, v.Name, v.Lat, v.Long)
		}
	}
}

func handlerName(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}
