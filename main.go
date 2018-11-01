package main

import (
	"encoding/xml"
	"log"
	"net/http"

	// "fmt"
	"io/ioutil"
	"os"

	"github.com/Athulus/feeds"
	"github.com/kr/pretty"
	"github.com/gorilla/mux"
)

func main() {

	f, err := os.Open("oneshotpodcast.rss")
	if err != nil {
		panic("AHH file bad")
	}
	feed := readRss(f)

	pretty.Println(feed)

	router := mux.NewRouter()
	router.HandleFunc("/serve/{file}", serveFile)
	router.HandleFunc("/serve", serveFile)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	name := "test.wav"
	vars := mux.Vars(r)
	category, exists := vars["file"]
	if exists {
		name = category
	}
	http.ServeFile(w, r, name)
}

func readRss(xmlFile *os.File) *feeds.RssFeed {
	var xmlFeed feeds.RssFeedXml
	bytes, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(bytes, &xmlFeed)

	return xmlFeed.Channel

}
