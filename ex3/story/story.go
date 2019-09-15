package story 

import (
	"os"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

func GetStory() map[string]Info {
	//first change working directory so it finds the file 
	os.Chdir("../story")

	//open json file 
	jsonFile, err := os.Open("story.json")

	if err != nil {
		fmt.Println(err)
		return nil
	}

	//otherwise the file was opened successfully so defer the closing of it 
	defer jsonFile.Close()
	defer os.Chdir("../main")

	//now get byte slice from opened file 
	byteVals, readErr := ioutil.ReadAll(jsonFile)
	if readErr != nil {
		fmt.Println(readErr)
		return nil
	}
	fmt.Println("Finished reading story")

	//var sm StoryMap 
	sm := make(map[string]Info) 

	json.Unmarshal(byteVals, &sm)

	return sm
}

type Info struct {
	Title string `json:"title"`
	Story []string `json:"story"`
	Options []OptObjects `json:"options"`
}

type OptObjects struct {
	Text string `json:"text"`
	Arc string `json:"arc"`
}