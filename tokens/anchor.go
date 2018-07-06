package tokens

import (
	"golang.org/x/net/html"
	"io"
)

// GetAnchorTagTokens returns all anchors tag tokens
func GetAnchorTagTokens(body io.Reader) ([]html.Token, error) {
	tokenizer := html.NewTokenizer(body)
	tokens := make([]html.Token, 10)

	for {
		token := tokenizer.Next()
		switch token {
		case html.ErrorToken:
			// If io.EOF is returned, then successful
			return tokens, tokenizer.Err()
		case html.StartTagToken:
			startTag := tokenizer.Token().Data
			if startTag == "a" {
				tokens = append(tokens, tokenizer.Token())
			}
		}
	}
}
