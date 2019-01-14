package client

import (
	"fmt"
	"log"

	"github.com/dghubble/go-twitter/twitter"
)

func (c *Client) ArchiveIDs(ids []int64) error {
	for _, id := range ids {
		log.Printf("[DEBUG] Older ID: %d", id)
		tweet, _, err := c.twitter.Statuses.Show(id, &twitter.StatusShowParams{})
		if err != nil {
			log.Printf("[DEBUG] Status %d was not found", id)
			continue
		}

		if !c.shouldArchiveTweet(*tweet) {
			log.Printf("[DEBUG] Tweet %d is too new - skipping!", tweet.ID)
			continue
		}

		err = c.delete(*tweet)
		if err != nil {
			return fmt.Errorf("Error deleting older tweet: %s", err)
		}
	}

	return nil
}
