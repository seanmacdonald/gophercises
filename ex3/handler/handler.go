package handler

import (
	"fmt"
	"os"
	"net/http"
	"html/template"
	"github.com/seanmacdonald/gophercises/ex3/story"
)

func MakeHandler(storyMap map[string]story.Info) (http.HandlerFunc, error) {
	fmt.Println("making handler")

	var hf http.HandlerFunc

	hf = func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[1:] //remove first char because its a slash 
		if info, ok := storyMap[path]; ok {
			renderChapter(info, w)
		} else {
			fmt.Fprintln(w, "nothing for: " + path)
		}
	}

	return hf, nil
}

func renderChapter(info story.Info, w http.ResponseWriter) {
	os.Chdir("../handler")
	defer os.Chdir("../main")

	tmpl, err := template.ParseFiles("layout.html")
	if err != nil {
		fmt.Println(err)
	}

	tmpl.Execute(w, info)
}