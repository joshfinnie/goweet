// Copyright 2014 Josh Finnie. All rights reserved.
//
// Use of this source code is governed by the MIT License license that can be
// found in the LICENSE file or at http://www.joshfinnie.com/license.txt

// This package Goweet takes an input of a Twitter handle and returns that
// person's last tweet.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

// returns the current implementation version
func Version() string {
	return "0.0.1"
}

// Structure for our Config
type Config struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

// Builds our config struct
func getConfig() Config {

	file, _ := os.Open("config.json")
	contents, _ := ioutil.ReadAll(file)

	var config Config
	json.Unmarshal(contents, &config)

	return config
}

// main application
func main() {
	config := getConfig()
	anaconda.SetConsumerKey(config.ConsumerKey)
	anaconda.SetConsumerSecret(config.ConsumerSecret)
	api := anaconda.NewTwitterApi(config.AccessToken, config.AccessSecret)
	arg := os.Args[1:]
	v := url.Values{}
	v.Set("from", string(arg[0]))
	v.Set("count", "1")
	searchResult, _ := api.GetSearch("", v)
	for _, tweet := range searchResult {
		fmt.Println(tweet.Text)
	}
}
