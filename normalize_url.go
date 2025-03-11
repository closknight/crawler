package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(URL string) (string, error) {
	u, err := url.Parse(URL)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %v", err)
	}

	path := u.Host + u.Path
	path = strings.ToLower(path)
	path = strings.TrimSuffix(path, "/")
	return path, nil
}
