package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	flag.Parse()

	fmt.Println(*urlFlag)

	response, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	io.Copy(os.Stdout, response.Body)
}
