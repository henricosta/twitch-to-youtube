package twitchtoyoutube

import (
	"os"
	"testing"
)

var TEST_CLIP_URL = "https://www.twitch.tv/cellbit/clip/AwkwardGlutenFreeMuleStoneLightning-KiUZC6pWzL1DzOq0"

func TestDownload(t *testing.T) {
	filepath, _, err := Download(TEST_CLIP_URL, "./videos")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(filepath); err != nil {
		t.Fatal(err)
	}
}
