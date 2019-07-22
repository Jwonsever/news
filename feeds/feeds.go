package feeds

import (
	"fmt"
	"os"
	"sort"
	"time"

	gfeed "github.com/gorilla/feeds"
	"github.com/mmcdole/gofeed"
)

var feeds = []string{
	"https://www.joelonsoftware.com/feed?numItems=100",
	//"https://blog.codinghorror.com/rss/", Giving bad XML...
	"https://www.ribice.ba/index.xml",
	"https://news.ycombinator.com/rss",
	"https://www.reddit.com/r/programming.rss",
}

type sortItems []*gfeed.Item

func (a sortItems) Len() int           { return len(a) }
func (a sortItems) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortItems) Less(i, j int) bool { return a[i].Created.Unix() < a[j].Created.Unix() }

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

	unsorted := []*gfeed.Item{}

	// Collect all the items.
	for _, feedURL := range feeds {
		if debug, _ := os.LookupEnv("debug"); debug == "1" {
			fmt.Println("Reading " + feedURL)
		}

		fp := gofeed.NewParser()
		feed, err := fp.ParseURL(feedURL)

		if err != nil {
			return "", err
		}

		for _, item := range feed.Items {
			unsorted = append(unsorted, convert(*item))
		}
	}

	//Sort the items by dates
	sort.Sort(sortItems(unsorted))
	combofeed.Items = unsorted

	//Return as feed
	return combofeed.ToRss()
}

// convert between feed-reader and feed-writer types.
func convert(i gofeed.Item) *gfeed.Item {
	if i.Author == nil {
		i.Author = &gofeed.Person{
			Email: "missing@email.com",
			Name:  "Author Tag Missing",
		}

	}

	if debug, _ := os.LookupEnv("debug"); debug == "1" {
		fmt.Printf("\n%v\n", i)
	}

	pub := i.PublishedParsed
	upd := i.UpdatedParsed
	if pub == nil {
		p := time.Now()
		pub = &p
	}
	if upd == nil {
		u := time.Now()
		upd = &u
	}

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
		Created: *pub,
		Updated: *upd,
		Content: i.Content,
	}
}
