package main

import (
	"fmt"
	"os"
	"testing"
)

var TEST_CLIP_URL = "https://www.twitch.tv/dlinkotoz/clip/FragileObliviousRutabagaHassanChop-RH_YBSPgDdfpawG5"

func TestDownload(t *testing.T) {
	t.Run("test if the function is downloading a clip", func(t *testing.T) {
		filepath := "test.mp4"
		err := Download(TEST_CLIP_URL, filepath)

		if err != nil {
			t.Errorf("Error in Download function")
		}

		_, err = os.Stat(filepath)
		if os.IsNotExist(err) {
			t.Errorf("File not found in directory")
		}
	})
}

func TestClipAuthenticatedUrl(t *testing.T) {
	t.Run("test if the function is returning a url for a valid clip", func(t *testing.T) {
		slug, _ := parseClipSlug(TEST_CLIP_URL)
		url := getClipAuthenticatedUrl(slug)

		if url == "" {
			t.Errorf("Error in ClipAuthenticatedUrl function")
		} else {
			fmt.Println(url)
		}
	})
}
