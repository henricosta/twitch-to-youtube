package twitch

import (
	"net/url"
	"strings"
)

func parseClipSlug(clipUrl string) (string, error) {
	u, err := url.Parse(clipUrl)

	if err != nil {
		return "", err
	}

	segments := strings.Split(u.Path, "/")
	slug := segments[len(segments)-1]

	return slug, err
}
