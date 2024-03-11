package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
)

func main() {

	fmt.Println("=====")
	fmt.Println("=====")
	fmt.Println("Map Read:")
	fmt.Println()
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	for path, url := range pathsToUrls {
		fmt.Printf("%s -> %s\n", path, url)
	}

	fmt.Println("=====")
	fmt.Println("=====")
	fmt.Println("JSON as Map Read:")
	fmt.Println()
	// Sample JSON {"name":"John", "age":30, "car":null}
	data := `{"name":"John", "age":30, "car":null}`
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(data), &m); err != nil {
		panic(err)
	}
	fmt.Println(m)

	fmt.Println("=====")
	fmt.Println("=====")
	fmt.Println("YAML Unmarshal:")
	fmt.Println()

	// type Pair struct {
	// 	Key   string
	// 	Value interface{}
	// }
	// var pairs []Pair
	// if err := json.Unmarshal([]byte(data), &pairs); err != nil {
	// 	panic(err)
	// }
	// fmt.Println(pairs)

	// type Order struct
	// type P struct {
	// 	Alias []string
	// 	RedUrl   []string
	// }

	// var pathsToUrls2 []P
	// if err := json.Unmarshal([]byte(data), &pathsToUrls2); err != nil {
	// 	panic(err)
	// }
	// fmt.Println(pathsToUrls2)

	// type Order struct {
	// 	CustomerName
	// var o Order

	yml := `
    - path: /urlshort
      url: https://github.com/gophercises/urlshort
    - path: /urlshort-final
      url: https://github.com/gophercises/urlshort/tree/solution
    `
	byml := []byte(yml)
	type Hodl struct {
		Path string `yaml:"path"`
		Url  string `yaml:"url"`
	}
	type Hodler []Hodl

	var h Hodler
	err := yaml.Unmarshal(byml, &h)
	if err == nil {
		fmt.Println("we did it fam")
		fmt.Println(h)
		fmt.Println("=====")
		for path, _ := range h {
			fmt.Println(h[path].Url)
		}
		//fmt.Println(h[1].Url)
	} else {
		fmt.Println(err)
	}

	// yamlHandler, err := handlers.YAMLHandler([]byte(yaml), mapHandler)
	// if err != nil {
	// 	panic(err)
	// }
	// func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {

	fmt.Println("=====")
	fmt.Println("=====")
	fmt.Println("YAML as byte array read/print:")
	fmt.Println()
	for _, value := range yml {
		fmt.Printf("%s", string(value))
	}

}
