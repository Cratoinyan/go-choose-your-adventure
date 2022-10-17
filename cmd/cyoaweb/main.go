package main

import (
	"cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fPath := flag.String("file", "story.json", "JSON file containing the story")
	port := flag.Int("port", 3000, "port to start the server on")
	flag.Parse()

	file, err := os.Open(*fPath)

	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(file)

	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story)

	fmt.Printf("Starting the server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
