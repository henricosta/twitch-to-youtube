package twitchtoyoutube

import (
	"fmt"
	"log"
	"os"

	"google.golang.org/api/youtube/v3"
)

func UploadVideo(filepath string, video *youtube.Video) (videoId string) {
	service := getService()

	call := service.Videos.Insert([]string{"snippet", "status"}, video)

	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		log.Fatalf("Error opening %v: %v", filepath, err)
	}

	fmt.Println("Uploading video...")
	response, err := call.Media(file).Do()
	if err != nil {
		log.Fatalf("Error making API call: %v", err.Error())
	}

	return response.Id
}
