package main 

import (
	"fmt"
	"net/http"
	"flag"
	"github.com/seanmacdonald/gophercises/ex3/story"
	"github.com/seanmacdonald/gophercises/ex3/handler"
)

func main() {
	var port = flag.String("port", "8080", "The port number that the server will use" )
	flag.Parse()
	*port = ":" + *port
	url := "http://localhost" + *port + "/"

	//retrieve the story
	storyMap := story.GetStory()

	//add links to the map 
	addLinks(storyMap, url)

	//verify that the story has an intro key
	if _, ok := storyMap["intro"]; !ok {
		fmt.Println("No story intro was found")
		return
	}

	//setup all story links on server
	storyHandler, err := handler.MakeHandler(storyMap)
	if err != nil {
		fmt.Println(err)
		return 
	}

	fmt.Println("Starting the server on " + *port)
	http.ListenAndServe(*port, storyHandler)
}

func addLinks(sm map[string]story.Info, url string) {
	for i, _ := range sm {
		for j, _ := range sm[i].Options {
			sm[i].Options[j].Url = url + sm[i].Options[j].Arc
		}
	}
}