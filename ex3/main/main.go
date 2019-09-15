package main 

import (
	"fmt"
	"net/http"
	"github.com/seanmacdonald/gophercises/ex3/story"
	"github.com/seanmacdonald/gophercises/ex3/handler"
)

func main() {
	fmt.Println("Starting story")

	//retrieve the story
	storyMap := story.GetStory()

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

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", storyHandler)
}