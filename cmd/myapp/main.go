package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	url := flag.String("url", "", "Download URL")
	fileName := flag.String("file", "file", "Location where file to be saved")

	flag.Parse()

	if *url == "" {
		fmt.Println("Error: Please provide a URL using flag -url")
		return
	}

	fmt.Println("Received URL:", *url)

	err := DownloadFile(fileName, url)
	if err != nil {
		fmt.Println("Error while downloading file", err)
	} else {
		fmt.Println("Download completed:", fileName)
	}
}

func DownloadFile(fileName, url *string) error {
	resp, err := http.Get(*url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(*fileName)
	if err != nil {
		fmt.Print("Error: Error while creating file")
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err

}
