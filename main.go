package main

import (
	"flag"
	"fmt"
	link "main/linkParser"
	"net/http"
	"net/url"
	"strings"
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

	reqUrl := response.Request.URL
	fmt.Print("reqUrl: ", reqUrl)

	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}

	base := baseUrl.String()

	var hrefs []string
	links, _ := link.Parse(response.Body)
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		}
	}

	for _, href := range hrefs {
		fmt.Println(href)
	}

}
