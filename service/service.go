package service

import (
	"fmt"
	"net/http"

	"github.com/vermagav/next-bart/config"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/name", handlerName)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, web service!")
	fmt.Fprintf(w, "\nStationsUrl: " + config.Bart.StationsUrl)
	fmt.Fprintf(w, "\nStationsUrl: " + config.Bart.EtdUrl)
}

func handlerName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, web service!\nYour name is: ")
	name := r.URL.Query().Get("name")
	if len(name) != 0 {
		w.Write([]byte(name))
	}
}
