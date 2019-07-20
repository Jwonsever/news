package feeds

import (
	"errors"
	"fmt"

	"github.com/mmcdole/gofeed"
)

var feeds = []string{
	"https://www.joelonsoftware.com/feed/",
}

func generateCombinedFeed() (string, error) {

	for _, feedUrl := range feeds {
		fmt.Println("Reading " + feedUrl)

		fp := gofeed.NewParser()
		feed, err := fp.ParseURL("http://feeds.twit.tv/twit.xml")

		if err != nil {
			return "", err
		}

		return feed.String(), nil
	}

	return "", errors.New("Failed to read feeds.")
}
