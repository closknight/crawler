package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("no website provided")
	} else if len(os.Args) > 2 {
		log.Fatalf("too many arguments provided")
	}
	baseURL := os.Args[1]
	fmt.Printf("starting crawl of: %s\n", baseURL)

	htmlBody, err := getHTML(baseURL)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(htmlBody)
}
