package client

import (
	"fmt"
	"time"

	"github.com/dghubble/go-twitter/twitter"
)

type Client struct {
	maximumAge    time.Duration
	tweetsPerPage int
	twitter       *twitter.Client
	username      string
	user          *twitter.User
}

func New(c *twitter.Client, username string) (*Client, error) {
	user, _, err := c.Users.Show(&twitter.UserShowParams{
		ScreenName: username,
	})
	if err != nil {
		return nil, fmt.Errorf("Error retrieving user: %s", err)
	}

	client := Client{
		maximumAge:    -7 * (24 * time.Hour),
		tweetsPerPage: 200,
		twitter:       c,
		username:      username,
		user:          user,
	}
	return &client, nil
}
