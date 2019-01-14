package client

import (
	"fmt"
	"log"

	"github.com/dghubble/go-twitter/twitter"
)

func (c *Client) ArchiveFavourites() error {
	var lastId *int64
	for true {
		params := &twitter.FavoriteListParams{
			UserID: c.user.ID,
			Count:  c.tweetsPerPage,
		}
		if lastId != nil {
			params.MaxID = (*lastId) - 1
		}
		favs, _, err := c.twitter.Favorites.List(params)
		if err != nil {
			return fmt.Errorf("Error retrieving favourites: %s", err)
		}
		log.Printf("Retrieved %d tweets..", len(favs))

		for _, v := range favs {
			if !c.shouldArchiveTweet(v) {
				log.Printf("[DEBUG] Tweet %d is too new, skipping!", v.ID)
				continue
			}

			err := c.unfavourite(v)
			if err != nil {
				return err
			}
		}

		newLastId := findOldestTweetId(favs)
		log.Printf("New Oldest ID is %d", newLastId)
		if newLastId == nil || newLastId == lastId {
			break
		}
		lastId = newLastId
	}
	return nil
}

func (c *Client) unfavourite(tweet twitter.Tweet) error {
	log.Printf("[DEBUG] Unfavouriting Tweet %d ", tweet.ID)
	_, _, err := c.twitter.Favorites.Destroy(&twitter.FavoriteDestroyParams{
		ID: tweet.ID,
	})
	if err != nil {
		return fmt.Errorf("Error unfavouriting %d: %s", tweet.ID, err)
	}
	log.Printf("         └ unfavourited %d ✅", tweet.ID)
	return nil
}
