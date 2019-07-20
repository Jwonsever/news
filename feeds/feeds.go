package feeds

import (
	"fmt"
	"time"

	gfeed "github.com/gorilla/feeds"
	"github.com/mmcdole/gofeed"
)

var feeds = []string{
	"https://www.joelonsoftware.com/feed?numItems=100",
}

// GenerateCombinedFeed generates a feed.
func GenerateCombinedFeed() (string, error) {
	now := time.Now()
	combofeed := gfeed.Feed{
		Title:       "News Aggregator",
		Link:        &gfeed.Link{Href: "http://jwonsever.com/feed"},
		Description: "Combined feed of many major software blogs.",
		Author:      &gfeed.Author{Name: "James Wonsever", Email: "jwonsever@gmail.com"},
		Created:     now,
	}
	combofeed.Items = []*gfeed.Item{}

	for _, feedURL := range feeds {
		fmt.Println("Reading " + feedURL)

		fp := gofeed.NewParser()
		feed, err := fp.ParseURL(feedURL)

		if err != nil {
			return "", err
		}

		for _, item := range feed.Items {
			combofeed.Items = append(combofeed.Items, convert(*item))
		}
	}

	//Shuffle items
	return combofeed.ToRss()
}

func convert(i gofeed.Item) *gfeed.Item {
	fmt.Println(i.Link)
	return &gfeed.Item{
		Title: i.Title,
		Link: &gfeed.Link{
			Href: i.Link,
		},
		Description: i.Description,
		Author: &gfeed.Author{
			Name:  i.Author.Name,
			Email: i.Author.Email,
		},
		Created: *i.PublishedParsed,
		Content: i.Content,
	}
}
