package main

import (
	"./tags"
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	c := &http.Client{}
	resp, err := c.Get("http://www.google.com")
	if err != nil {
		panic(err)
	}

	lists := []string{"1", "2", "3", "4"}

	for _, text := range lists {
		var tokens []html.Token
		var atags []byte

		switch text {
		case "1":
			fmt.Println("Reached")
			tokens, err = tags.GetAnchorTagTokens(resp.Body)
		case "2":
			atags, err = tags.GetAnchorTagTexts(resp.Body)
		case "3":
			atags, err = tags.GetAnchorTagNames(resp.Body)
		case "4":
			atags, err = tags.GetAnchorTagAttrs(resp.Body)
		}

		if atags != nil {
			for _, tag := range atags {
				fmt.Println(string(tag))
			}
		} else {
			for _, token := range tokens {
				fmt.Println(token)
			}
		}
	}

	defer resp.Body.Close()
}
