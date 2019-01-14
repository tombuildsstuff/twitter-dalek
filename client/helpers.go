package client

import (
	"log"
	"time"

	"github.com/dghubble/go-twitter/twitter"
)

func findOldestTweetId(input []twitter.Tweet) *int64 {
	if len(input) == 0 {
		return nil
	}

	oldestId := int64(-1)
	oldestPublishDate := time.Now()
	for _, v := range input {
		t, _ := v.CreatedAtTime()
		if t.Before(oldestPublishDate) {
			oldestId = v.ID
			log.Printf("[DEBUG] New Assumed Oldest ID is %d", v.ID)
		}
	}

	return &oldestId
}

func (c *Client) shouldArchiveTweet(tweet twitter.Tweet) bool {
	threshold := time.Now().Add(c.maximumAge)
	time, _ := tweet.CreatedAtTime()
	return time.Before(threshold)
}
