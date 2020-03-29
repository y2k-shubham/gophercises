package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/y2k-shubham/gophercises-y2k/urlshort"
)

func main() {
	mapHandler := urlshort.BuildDefaultHandler()

	pathRulesString, isYml := GetPathRules()
	usedHandler, err := urlshort.UnmarshalHandler([]byte(pathRulesString), isYml, mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	//http.ListenAndServe(":8080", mapHandler)
	http.ListenAndServe(":8080", usedHandler)
}

func GetPathRules() (string, bool) {
	var pathRulesString string
	var isYml bool
	// go run urlshort/main/main.go -yaml-file my-rules.yml
	yamlFileName := flag.String("yaml-file", "", "YAML file containing path rules")
	// go run urlshort/main/main.go -json-file my-rules.json
	jsonFileName := flag.String("json-file", "", "JSON file containing path rules")

	flag.Parse()

	if *jsonFileName != "" {
		pathRulesString = urlshort.ReadFile(*jsonFileName)
		isYml = false
	} else {
		isYml = true
		if *yamlFileName != "" {
			pathRulesString = urlshort.ReadFile(*yamlFileName)
		} else {
			pathRulesString = GetDefaultYamlRules()
		}
	}

	return pathRulesString, isYml
}

func GetDefaultYamlRules() string {
	return `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
}
