package text

import (
	"io"

	"golang.org/x/net/html"
)

// GetAnchorTagTexts returns all anchors tag text
func GetAnchorTagTexts(body io.Reader) ([]byte, error) {
	tokenizer := html.NewTokenizer(body)
	tagTexts := make([]byte, 10)

	for {
		token := tokenizer.Next()
		switch token {
		case html.ErrorToken:
			// If io.EOF is returned, then successful
			return tagTexts, tokenizer.Err()
		case html.StartTagToken:
			startTag := tokenizer.Token().Data
			if startTag == "a" {
				tagTexts = append(tagTexts, tokenizer.Text()...)
			}
		}
	}
}
