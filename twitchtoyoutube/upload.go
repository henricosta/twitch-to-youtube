package twitchtoyoutube

import (
	"log"
	"os"

	"google.golang.org/api/youtube/v3"
)

func UploadVideo(filepath string, video *youtube.Video) (videoId string) {
	service := getService()

	call := service.Videos.Insert([]string{"snippet", "status"}, video)

	file, err := os.Open("test.mp4")
	defer file.Close()
	if err != nil {
		log.Fatalf("Error opening %v: %v", "test.mp4", err)
	}

	response, err := call.Media(file).Do()
	if err != nil {
		log.Fatalf("Error making API call: %v", err.Error())
	}

	return response.Id
}
