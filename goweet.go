// Copyright 2014 Josh Finnie. All rights reserved.
//
// Use of this source code is governed by the MIT License license that can be
// found in the LICENSE file or at http://www.joshfinnie.com/license.txt

// This package Goweet takes an input of a Twitter handle and returns that
// person's last tweet.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mrjones/oauth"
)

// returns the current implementation version
func Version() string {
	return "0.0.1"
}

const tweetURL string = "https://api.twitter.com/1.1/statuses/user_timeline.json?screen_name=XXXXX&count=1"

// main application
func main() {
	c := oauth.NewConsumer(
		"US55QyDCABM9jjOYG4q4IThde",
		"xujBz77BlHgkQ0L8ns9WPPiSCL2ixGBd5nx2Timlo8hMO2ml11",
		oauth.ServiceProvider{
			RequestTokenUrl:   "http://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		},
	)
	arg := os.Args[1:]
	fmt.Println(len(arg))
	fmt.Println(strings.Replace(tweetURL, "XXXXX", arg[0], 1))
	response, err := c.Get(
		strings.Replace(tweetURL, "XXXXX", arg[0], 1),
		map[string]string{"count": "1"},
		"8820492-ru3XP6qzsDIevjbg0YbU2OeBw3kj6P431ixesFhSOG")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	bits, err := ioutil.ReadAll(response.Body)
	fmt.Println("The newest item in your home timeline is: " + string(bits))
}
