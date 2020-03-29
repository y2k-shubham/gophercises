package urlshort

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"log"
	"net/http"
)

func BuildDefaultHandler() http.HandlerFunc {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := MapHandler(pathsToUrls, mux)
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

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if path, ok := pathsToUrls[request.URL.Path]; ok {
			http.Redirect(writer, request, path, http.StatusFound)
		} else {
			fallback.ServeHTTP(writer, request)
		}
	}
}

// UnmarshalHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func UnmarshalHandler(bytes []byte, isYml bool, fallback http.Handler) (http.HandlerFunc, error) {
	pathMap := make(map[string]string)
	var err error

	// had to duplicate code just because type in list of maps 'm' differs for yaml & json unmarshal methods
	if isYml {
		var m []map[interface{}]interface{}
		err = yaml.Unmarshal(bytes, &m)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		for _, myMap := range m {
			pathMap[myMap["path"].(string)] = myMap["url"].(string)
		}
	} else {
		var m []map[string]interface{}
		err = yaml.Unmarshal(bytes, &m)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		for _, myMap := range m {
			pathMap[myMap["path"].(string)] = myMap["url"].(string)
		}
	}

	return MapHandler(pathMap, fallback), err
}
