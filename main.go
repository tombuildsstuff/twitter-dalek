package main

import (
	"log"
	"os"

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

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	twitterClient := twitter.NewClient(httpClient)

	log.Printf("[DEBUG] Building Twitter Client..")
	client, err := client.New(twitterClient, username)
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
