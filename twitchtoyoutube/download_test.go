package twitchtoyoutube

import "testing"

var TEST_CLIP_URL = "https://www.twitch.tv/cellbit/clip/AwkwardGlutenFreeMuleStoneLightning-KiUZC6pWzL1DzOq0"

func TestDownload(t *testing.T) {
	err := Download(TEST_CLIP_URL)
	if err != nil {
		t.Fatal(err)
	}
}
