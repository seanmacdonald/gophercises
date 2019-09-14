package main 

import (
	"fmt"
	//"html/template"
	//"net/http"
	//"encoding/json"
	"github.com/seanmacdonald/gophercises/ex3/story"
)

func main() {
	fmt.Println("Starting story")

	// step 1 - parse the story from the json 
	storyMap := story.GetStory()

	//check if all the info was parse correctly 
	for i, val := range storyMap {
		fmt.Println(i)
		fmt.Println(val)
	} 

	//step 2 - serve correct page using html/template 

}