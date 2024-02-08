package twitch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	client *http.Client
	gqlUrl string
}

type Clip struct {
	Title string `json:"title"`
}

func parseClipSlug(clipUrl string) (string, error) {
	u, err := url.Parse(clipUrl)

	if err != nil {
		return "", err
	}

	segments := strings.Split(u.Path, "/")
	slug := segments[len(segments)-1]

	return slug, err
}

func authenticatedPost(c *Client, url string, query string) ([]byte, error) {
	bodyBytes, err := json.Marshal(map[string]string{"query": query})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Client-ID", "kimne78kx3ncx6brgo4mv6wki5h1ko")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func getClip(client *Client, slug string) (Clip, error) {
	query := fmt.Sprintf(`{clip(slug: "%s"){title}}`, slug)

	reponseBody, err := authenticatedPost(client, client.gqlUrl, query)
	if err != nil {
		return Clip{}, err
	}

	type response struct {
		Data struct {
			Clip `json:"clip"`
		} `json:"data"`
	}
	var r response
	err = json.Unmarshal(reponseBody, &r)
	if err != nil {
		return Clip{}, err
	}

	return r.Data.Clip, nil
}
