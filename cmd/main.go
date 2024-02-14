package main

import (
	"fmt"
	"log"
	"os"

	twyt "github.com/henricosta/twitch-to-youtube/twitchtoyoutube"
	"google.golang.org/api/youtube/v3"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a url as an argument.")
		return
	}

	url := os.Args[1]

	filepath, err := twyt.Download(url)
	if err != nil {
		log.Fatal(err)
	}

	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       "Test Upload",
			Description: "Test Description",
		},
	}

	twyt.UploadVideo(filepath, upload)

	fmt.Println("Video uploaded successfully!")
}
