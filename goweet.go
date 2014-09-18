// Copyright 2014 Josh Finnie. All rights reserved.
//
// Use of this source code is governed by the MIT License license that can be
// found in the LICENSE file or at http://www.joshfinnie.com/license.txt

// This package Goweet takes an input of a Twitter handle and returns that
// person's last tweet.
package main

import (
	"fmt"
    	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

// returns the current implementation version
func Version() string {
	return "0.0.1"
}


// main application
func main() {
    anaconda.SetConsumerKey("US55QyDCABM9jjOYG4q4IThde")
    anaconda.SetConsumerSecret("xujBz77BlHgkQ0L8ns9WPPiSCL2ixGBd5nx2Timlo8hMO2ml11")
    api := anaconda.NewTwitterApi("8820492-ru3XP6qzsDIevjbg0YbU2OeBw3kj6P431ixesFhSOG", "Aqh0fmAlZeFUlRWkxzjizkEAieLsCGfkFhKZqgqCpcmBb")
	arg := os.Args[1:]
    v := url.Values{}
    v.Set("from", string(arg[0]))
    v.Set("count", "1")
    searchResult, _ := api.GetSearch("", v)
    for _ , tweet := range searchResult {
        fmt.Println(tweet.Text)
    }   
}
