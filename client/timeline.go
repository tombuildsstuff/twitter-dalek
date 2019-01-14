package client

import (
	"fmt"
	"log"

	"github.com/dghubble/go-twitter/twitter"
)

func (c *Client) ArchiveTimeline(includeRetweets bool) error {
	var lastId *int64
	for true {
		params := twitter.UserTimelineParams{
			UserID:          c.user.ID,
			IncludeRetweets: &includeRetweets,
			Count:           c.tweetsPerPage,
		}
		if lastId != nil {
			params.MaxID = (*lastId) - 1
		}
		retrieved, _, err := c.twitter.Timelines.UserTimeline(&params)
		if err != nil {
			return err
		}

		log.Printf("Retrieved %d tweets..", len(retrieved))
		for _, v := range retrieved {
			if !c.shouldArchiveTweet(v) {
				log.Printf("[DEBUG] Tweet %d is too new, skipping!", v.ID)
				continue
			}

			log.Printf(v.Text)
			err := c.delete(v)
			if err != nil {
				return fmt.Errorf("Error deleting %d: %s", v.ID, err)
			}
		}

		newLastId := findOldestTweetId(retrieved)
		log.Printf("New Oldest ID is %d", newLastId)
		if newLastId == nil || newLastId == lastId {
			break
		}
		lastId = newLastId
	}

	return nil
}

func (c *Client) delete(tweet twitter.Tweet) error {
	if tweet.Retweeted {
		log.Printf("Unretweeting %d", tweet.ID)
		_, _, err := c.twitter.Statuses.Unretweet(tweet.ID, &twitter.StatusUnretweetParams{
			ID: tweet.ID,
		})
		if err != nil {
			log.Printf("Error unretweeting %d", tweet.ID)
		}
		return nil
	}

	log.Printf("Deleting %d", tweet.ID)
	_, _, err := c.twitter.Statuses.Destroy(tweet.ID, &twitter.StatusDestroyParams{
		ID: tweet.ID,
	})
	if err != nil {
		log.Printf("Error deleting %d", tweet.ID)
	}
	return nil
}
