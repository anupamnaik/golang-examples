package functions

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// SiteMapIndex the stitemap xml element
type SiteMapIndex struct {
	Locations []Location `xml:"sitemap"`
}

// Location the location xml element
type Location struct {
	Loc string `xml:"loc"`
}

// Parse parse the xml and print
func Parse() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var s SiteMapIndex
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Locations)
}
