package tags

import (
	"golang.org/x/net/html"
	"io"
)

// GetAnchorTagNames returns all anchors tag names
func GetAnchorTagNames(body io.Reader) ([]byte, error) {
	tokenizer := html.NewTokenizer(body)
	tagNames := make([]byte, 10)

	for {
		token := tokenizer.Next()
		switch token {
		case html.ErrorToken:
			// If io.EOF is returned, then successful
			return tagNames, tokenizer.Err()
		case html.StartTagToken:
			startTag := tokenizer.Token().Data
			if startTag == "a" {
				tagName, _ := tokenizer.TagName()
				tagNames = append(tagNames, tagName...)
			}
		}
	}
}

// GetAnchorTagAttrs returns all anchors tag text
func GetAnchorTagAttrs(body io.Reader) ([]byte, error) {
	tokenizer := html.NewTokenizer(body)
	tokens := make([]byte, 10)

	for {
		token := tokenizer.Next()
		switch token {
		case html.ErrorToken:
			// If io.EOF is returned, then successful
			return tokens, tokenizer.Err()
		case html.StartTagToken:
			startTag := tokenizer.Token().Data
			if startTag == "a" {
				attrKey, attrVal, _ := tokenizer.TagAttr()
				tokens = append(tokens, attrKey...)
				tokens = append(tokens, attrVal...)
			}
		}
	}
}
