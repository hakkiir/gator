package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	var reader io.Reader
	var RSS RSSFeed

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, reader)
	if err != nil {
		return nil, err
	}

	client := http.Client{Timeout: 10 * time.Second}

	req.Header.Set("User-Agent", "gator")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(body, &RSS)
	if err != nil {
		return nil, err
	}

	for i, item := range RSS.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		RSS.Channel.Item[i] = item
	}
	RSS.Channel.Description = html.UnescapeString(RSS.Channel.Description)
	RSS.Channel.Title = html.UnescapeString(RSS.Channel.Title)
	return &RSS, nil

}
