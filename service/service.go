package service

import (
	"fmt"
	"net/http"

	"github.com/vermagav/next-bart/bart"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/name", handlerName)
}

func handler(w http.ResponseWriter, r *http.Request) {
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

func handlerName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, web service!\nYour name is: ")
	name := r.URL.Query().Get("name")
	if len(name) != 0 {
		w.Write([]byte(name))
	}
}
