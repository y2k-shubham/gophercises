package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/y2k-shubham/gophercises-y2k/urlshort"
)

func main() {
	mapHandler := BuildDefaultHandler()

	var pathRulesString string
	// go run urlshort/main/main.go -yaml-file my-rules.yml
	yamlFileName := flag.String("yaml-file", "", "YAML input file containing path rules")

	flag.Parse()
	if *yamlFileName != "" {
		pathRulesString = urlshort.ReadYamlFile(*yamlFileName)
	} else {
		pathRulesString = GetDefaultYamlRules()
	}

	yamlHandler, err := urlshort.YAMLHandler([]byte(pathRulesString), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	//http.ListenAndServe(":8080", mapHandler)
	http.ListenAndServe(":8080", yamlHandler)
}

func GetDefaultYamlRules() string {
	return `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
}

func BuildDefaultHandler() http.HandlerFunc {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	return mapHandler
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
