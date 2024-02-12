package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Download(url string, filepath string) error {
	slug, err := parseClipSlug(url)
	if err != nil {
		return err
	}

	client := &Client{
		client: &http.Client{},
		gqlUrl: "https://gql.twitch.tv/gql",
	}

	clip, err := getClip(client, slug)
	if err != nil {
		return err
	}

	fmt.Printf("Found clip: %v\n", clip.Title)

	videoUrl := getClipAuthenticatedUrl(slug)

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(videoUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
