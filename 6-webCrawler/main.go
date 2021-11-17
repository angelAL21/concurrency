package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
	//"golang.org/x/net/html"
)

var fetched map[string]bool

type result struct {
	url   string
	urls  []string
	err   error
	depth int
}

//Crawl uses findlinks to recursively crawl.
func Crawl(url string, depth int) {
	ch := make(chan *result) //comm btween crawl and main

	fetch := func(url string, depth int) {
		urls, err := findLinks(url)

		ch <- &result{url, urls, err, depth}
	}

	go fetch(url, depth)
	fetched[url] = true

	for fetching := 1; fetching > 0; fetching-- {
		res := <-ch //passing the value to chan

		if res.err != nil { //checking is all ok.
			continue
		}
		fmt.Printf("found %s\n", res.url) //printing...
		if res.depth > 0 {
			for _, u := range res.urls {
				if !fetched[u] {
					fetching++
					go fetch(u, res.depth-1) //printing nex lvl of url.
					fetched[u] = true
				}
			}
		}
	}
	close(ch)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fetched = make(map[string]bool)
	now := time.Now()
	Crawl("https://youtube.com", 2)
	fmt.Println("time taken:", now)
}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

// visit appends to links each link found in n, and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
