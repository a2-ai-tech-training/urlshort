package handlers

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v3"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	fmt.Println("MapHandler called")
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if url, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, url, http.StatusFound)
			return
		} else {
			fallback.ServeHTTP(w, r)
		}

	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.

type Hodl struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

type Hodler []Hodl

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	fmt.Println("YAMLHandler called")
	return func(w http.ResponseWriter, r *http.Request) {
		var h Hodler
		err := yaml.Unmarshal([]byte(yml), &h)
		if err != nil {
			fmt.Errorf("yaml unmarshal error: %s", err)
			return
		}
		fmt.Printf("YAML Handler serving path: %s\n", r.URL.Path)

		for path := range h {

			if h[path].Path == r.URL.Path {
				fmt.Printf("Redirecting to %s\n", h[path].Url)
				http.Redirect(w, r, h[path].Url, http.StatusFound)
				return
			}
			continue
		}
		fallback.ServeHTTP(w, r)
	}, nil
}
