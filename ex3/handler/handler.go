package handler

import (
	"fmt"
	"net/http"
	"github.com/seanmacdonald/gophercises/ex3/story"
)

func MakeHandler(storyMap map[string]story.Info) (http.HandlerFunc, error) {
	fmt.Println("making handler")

	var hf http.HandlerFunc

	hf = func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[1:] //remove first char because its a slash 
		if info, ok := storyMap[path]; ok {
			//TODO: call html func to make it look nice 
			fmt.Fprintln(w, info.Title)
		} else {
			fmt.Fprintln(w, "nothing for: " + path)
		}
	}

	return hf, nil
}