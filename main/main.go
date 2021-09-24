package main

import (
	"fmt"
	"net/http"

	"github.com/UfiairENE/urlshortner"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/UfiairENE/urlshortner",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshortner.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /urlshortner
  url: https://github.com/UfiairENE/urlshortner
- path: /urlshortner-final
  url: https://github.com/UfiairENE/urlshortner/tree/solution
`
	yamlHandler, err := urlshortner.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Welcome)
	return mux
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome, Irene!")
}
