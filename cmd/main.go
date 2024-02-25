package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	twyt "github.com/henricosta/twitch-to-youtube/twitchtoyoutube"
	"google.golang.org/api/youtube/v3"
)

var url, title, privacy, dir string

func main() {
	flag.StringVar(&url, "url", "", "URL of the clip")
	flag.StringVar(&title, "title", "", "Title of the video")
	flag.StringVar(&privacy, "privacy", "private", "Privacy status of the video")
	flag.StringVar(&dir, "dir", "./videos", "Folder where downloaded clips will be saved")

	flag.Parse()

	if url == "" {
		fmt.Println("URL is required")
		os.Exit(1)
	}

	filepath, clipTitle, err := twyt.Download(url, dir)
	if err != nil {
		log.Fatal(err)
	}

	if title == "" {
		title = clipTitle
	}

	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       title,
			Description: "Test Description",
		},
		Status: &youtube.VideoStatus{
			PrivacyStatus: privacy,
		},
	}

	twyt.UploadVideo(filepath, upload)

	fmt.Println("Video uploaded successfully!")
}
