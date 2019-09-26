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
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

// returns the current implementation version
func Version() string {
	return "0.0.1"
}

// Config is the data structure for our config file
type Config struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

// LoadConfiguration builds our config struct from a json file.
func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config

}

// main application
func main() {
	config := LoadConfiguration("config.json")
	api := anaconda.NewTwitterApiWithCredentials(config.AccessToken, config.AccessSecret, config.ConsumerKey, config.ConsumerSecret)
	arg := os.Args[1:]
	v := url.Values{}
	v.Set("from", string(arg[0]))
	v.Set("count", "1")
	searchResult, _ := api.GetSearch("", v)
	for _, tweet := range searchResult.Statuses {
		fmt.Println(tweet.Text)
	}
}
