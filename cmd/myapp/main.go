package main

import (
	"flag"
	"fmt"
)

func main() {
	url := flag.String("url", "", "Download URL")

	flag.Parse()

	if *url == "" {
		fmt.Println("Error: Please provide a URL using flag -url")
		return
	}

	fmt.Println("Received URL:",*url)
}