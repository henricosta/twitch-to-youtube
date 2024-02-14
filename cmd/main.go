package main

import (
	"fmt"
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

	twyt.Download(url)

	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       "Test Upload",
			Description: "Test Description",
		},
	}

	twyt.UploadVideo("./clip.mp4", upload)

	fmt.Println("Video uploaded successfully!")
}
