package functions

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Crawl get and print the content of a site
func Crawl() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	stringBody := string(bytes)

	fmt.Println(stringBody)
	resp.Body.Close()
}
