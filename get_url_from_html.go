package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("could not parse base URL: %v", err)
	}

	reader := strings.NewReader(htmlBody)
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, fmt.Errorf("could not parse html: %v", err)
	}

	var links []string
	for node := range doc.Descendants() {
		if node.Type == html.ElementNode && node.DataAtom == atom.A {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					link, err := url.Parse(attr.Val)
					if err != nil {
						fmt.Printf("could not parse url %v: %v\n", attr.Val, err)
						continue
					}
					resolvedURL := baseURL.ResolveReference(link)
					links = append(links, resolvedURL.String())
				}
			}
		}
	}
	return links, nil
}
