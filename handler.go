package urlshort

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			fmt.Println("Redirecting to:", dest)
			return
		}

		fmt.Println("Loading:", path)
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(y []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(y, &pathUrls)
	if err != nil {
		return nil, err
	}
	pathsToUrls := make(map[string]string)

	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.URL
	} 
	return MapHandler(pathsToUrls, fallback), nil
}
