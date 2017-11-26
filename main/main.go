package main

import (
	"fmt"
	"net/http"

	"github.com/zacscodingclub/urlshort"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/url-short-godoc": "https://godoc.org/github.com/zacscodingclub/urlshort",
		"/yaml-godoc":      "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yaml := `
- path: /urlshort
  url: https://github.com/zacscodingclub/urlshort
- path: /urlshort-final
  url: https://github.com/zacscodingclub/urlshort/tree/solution
`

	// equivalent json
	//[
	//	{"path": "/value", "url": "https://example.com"},
	//	{"path": "/value", "url": "https://example.com"},
	//]

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	port := ":3000"
	fmt.Println("Started server on port", port)
	http.ListenAndServe(port, yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hey there world!")
}
