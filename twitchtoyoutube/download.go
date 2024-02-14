package twitchtoyoutube

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Download(url string) (string, error) {
	slug, err := parseClipSlug(url)
	if err != nil {
		return "", err
	}

	client := &Client{
		client: &http.Client{},
		gqlUrl: "https://gql.twitch.tv/gql",
	}

	clip, err := getClip(client, slug)
	if err != nil {
		return "", err
	}

	fmt.Printf("Found clip: %v\n", clip.Title)

	videoUrl := getClipAuthenticatedUrl(slug)

	dir := "./videos"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return "", err
		}
	}

	filepath := filepath.Join(dir, clip.Title+".mp4")
	out, err := os.Create(filepath)
	if err != nil {
		fmt.Println("error while creating directory")
		return "", err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(videoUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return filepath, err
}
