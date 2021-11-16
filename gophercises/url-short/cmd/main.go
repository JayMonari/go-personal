package main

import (
	"fmt"
	"net/http"
	"urlshort"
)

// TODO Update to accept a YAML file as a flag and load that in place of the
// hardcoded string.
// TODO Build a Handler that doesn't read from a map but instead reads from a
// database: BoltDB or Postgres.
func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
  jDoc := `[
{"path": "/json", "url": "https://www.w3schools.com/whatis/whatis_json.asp"},
{"path": "/429", "url": "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429"}
]`

  jsonHandler, err := urlshort.JSONHandler([]byte(jDoc), yamlHandler)
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
