package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	link "main/linkParser"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}
type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	maxDepth := flag.Int("depth", 3, "the maximum number of links deep to traverse")
	flag.Parse()

	pages := bfs(*urlFlag, *maxDepth)

	toXml := urlset{
		Xmlns: xmlns,
	}
	for _, page := range pages {
		toXml.Urls = append(toXml.Urls, loc{page})
	}

	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(toXml); err != nil {
		panic(err)
	}
	fmt.Println()
}

func bfs(urlStr string, maxDepth int) []string {
	seen := make(map[string]struct{})
	var q map[string]struct{}
	nq := map[string]struct{}{
		urlStr: struct{}{},
	}

	for i := 0; i <= maxDepth; i++ {
		q, nq = nq, make(map[string]struct{})
		if len(q) == 0 {
			break
		}
		for url, _ := range q {

			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = struct{}{}

			for _, link := range get(url) {
				if _, ok := seen[link]; !ok {
					nq[link] = struct{}{}
				}
				nq[link] = struct{}{}
			}
		}
	}

	ret := make([]string, 0, len(seen))
	for url, _ := range seen {
		ret = append(ret, url)
	}

	return ret
}

func get(urlStr string) []string {
	response, err := http.Get(urlStr)
	if err != nil {
		return []string{}
	}
	defer response.Body.Close()

	reqUrl := response.Request.URL

	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()

	return filter(base, hrefs(response.Body, base))
}

func hrefs(r io.Reader, base string) []string {

	var ret []string
	links, _ := link.Parse(r)
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}
	return ret
}

func filter(base string, links []string) []string {
	var ret []string
	for _, link := range links {
		if strings.HasPrefix(link, base) {
			ret = append(ret, link)
		}
	}
	return ret
}
