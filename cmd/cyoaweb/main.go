package main

import (
	"cyoa"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	fPath := flag.String("file", "story.json", "JSON file containing the story")
	flag.Parse()

	file, err := os.Open(*fPath)

	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(file)

	var story cyoa.Story

	if err := decoder.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
