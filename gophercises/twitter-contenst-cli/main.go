package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/oauth2"
)

func main() {
	key := os.Getenv("TWITTER_API_KEY")
	secret := os.Getenv("TWITTER_API_SECRET")
	req, err := http.NewRequest("POST", "https://api.twitter.com/oauth2/token", strings.NewReader("grant_type=client_credentials"))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(key, secret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	var c http.Client
	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var tok oauth2.Token
	if err := json.NewDecoder(res.Body).Decode(&tok); err != nil {
		log.Fatal(err)
	}
	tclient := (&oauth2.Config{}).Client(context.Background(), &tok)
	res2, err := tclient.Get("https://api.twitter.com/2/tweets/991053593250758658/retweeted_by")
	if err != nil {
		log.Fatal(err)
	}
	defer res2.Body.Close()

	io.Copy(os.Stdout, res2.Body)
}
