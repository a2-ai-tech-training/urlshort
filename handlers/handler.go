package handlers

import (
	//"fmt"
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
	//	TODO: Implement this...
	fmt.Println("MapHandler called")
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if url, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, url, http.StatusFound)
			return
		} else {
			fallback.ServeHTTP(w, r)
		}

		//return nil
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
			fmt.Errorf("yaml unmarshal error %w", err)
			return
		}
		fmt.Printf("YAML Handler serving path: %s\n", r.URL.Path)
		// if r.URL.Path != h[*].Path {
		// 	return
		// }
		//return w, r h[path].Url
		for path := range h {
			//fmt.Fprint(w, "hi its me")
			// if r.URL.Path == h[path].Path {
			// 	//r.URL.Path = h[path].Url
			// 	fmt.Fprintf(os.Stdout, "r path before: %s", r.URL.Path)
			// 	fmt.Fprintf(os.Stdout, "r path after: %s", h[path].Url)
			// 	http.Redirect(w, r, h[path].Url, http.StatusFound)
			// } else {
			// 	return
			// }

			//fmt.Printf("YAML Read %s\n", h[path].Path)
			//fmt.Printf("YAML Redirect URL: %s\n", h[path].Url)
			//http.Redirect(w, r, h[path].Url, http.StatusFound)
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

// 		// url := h[Path].Url{
// 		// 	http.Redirect(w, r, url, http.StatusFound),
// 		// 	//return url
// 		// }
// 		// return fallback.ServeHTTP(w, r)
// 	}, nil
// }

// func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {

// 	var h Hodler
// 	// TODO: Implement this...
// 	// TODO: Check out yaml unmarshalling
// 	err := yaml.Unmarshal([]byte(yml), &h)
// 	if err != nil {
// 		fmt.Printf("yaml Unmarshal failed, %s", err)
// 	}

// 	for path, _ := range h {
// 		// if _ == _ {
// 		// 	fmt.Println("Something went wrong. _ = 123")
// 		// }
// 		if h[path].Url == "" {
// 			return nil, nil
// 		} else {
// 			return func(w http.ResponseWriter, r *http.Request), nil{http.Redirect(w, r, url, http.StatusFound), nil}
// 		}
// 	}
// 	//return hello2, nil
// }

// func hello0(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Placeholder 0 read")
// }
// func hello1(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Placeholder 1 read")
// }
// func hello2(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Placeholder 2 read")
// }

// 		return func(w http.ResponseWriter, r *http.Request), nil {
// 			http.Redirect(w, r, url, http.StatusFound), nil,
// 		}
// 		} else {
// 	return nil, nil
// 	}
// }
// return func(w http.ResponseWriter, r *http.Request), nil {

// 	// path := r.URL.Path
// 	for path, url in range h {

// 	}
// 	url, ok := h[].URL: ok {
// 		http.Redirect(w, r, url, http.StatusFound), nil
// 		return
// 	} else {
// 		fallback.ServeHTTP(w, r)
// 	}
// }
// return nil, nil
