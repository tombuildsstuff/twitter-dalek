package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/tombuildsstuff/twitter-dalek/client"
)

func main() {
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret := os.Getenv("TWITTER_ACCESS_SECRET")
	username := os.Getenv("TWITTER_USERNAME")
	daysToKeepRaw := os.Getenv("DAYS_TO_KEEP")
	daysToKeep := 7
	var err error
	if daysToKeepRaw != "" {
		daysToKeep, err = strconv.Atoi(daysToKeepRaw)
		if err != nil {
			panic(fmt.Sprintf("DAYS_TO_KEEP should be an int but got %s", os.Getenv("DAYS_TO_KEEP")))
		}
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	twitterClient := twitter.NewClient(httpClient)

	log.Printf("[DEBUG] Building Twitter Client..")
	client, err := client.New(twitterClient, username, daysToKeep)
	if err != nil {
		panic(err)
	}

	log.Printf("[DEBUG] Archiving Timeline..")
	err = client.ArchiveTimeline(true)
	if err != nil {
		panic(err)
	}

	log.Printf("[DEBUG] Archiving Favourites..")
	err = client.ArchiveFavourites()
	if err != nil {
		panic(err)
	}
}
