package urlshort

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	var hf http.HandlerFunc

	hf = func(w http.ResponseWriter, r *http.Request) {
		var req_path string
		req_path = r.URL.String()

		if val, ok := pathsToUrls[req_path]; ok {
			fmt.Println("Go to this url: " + val)
			http.Redirect(w, r, val, http.StatusSeeOther)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}

	return hf
}

// YAMLHandler will parse the provided YAML and then return
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
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	urlInfoMap := parseYaml(yml)
	fmt.Println(urlInfoMap)
	
	//setup handler 
	var hf http.HandlerFunc
	hf = func (w http.ResponseWriter, r *http.Request) {
		fallback.ServeHTTP(w, r)
	}

	return hf, nil
}

type UrlInfo struct {
	Path string `yaml:"path"`
	Url string	`yaml:"url"`
}

type UrlInfoMap struct {
	Records []UrlInfo
}

// private function for parsing yaml 
// Here is a good resource for learning how to parse yaml with the library: 
// http://squarism.com/2014/10/13/yaml-go/ 
func parseYaml(yml []byte) map[string]string {
	var b UrlInfoMap

	err := yaml.Unmarshal(yml, &b)
	if err != nil {
		fmt.Println(err)
	}

	//go through object and creatre map[string]string 
	urlMap := make(map[string]string) 

	for _, val := range b.Records {
		urlMap[val.Path] = val.Url
	}

	return urlMap
}
