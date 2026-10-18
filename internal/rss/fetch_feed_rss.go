package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	
}


func FetchFeed(ctx context.Context , feedURL string) (*RSSFeed , error){

	req , err := http.NewRequestWithContext(ctx , "GET" , feedURL , nil)
	if err != nil {
		fmt.Println("Error sending a req")
		os.Exit(1)
	}

	req.Header.Set("User-Agent" , "gator")

	client := http.Client{}
	res , err := client.Do(req)
	if err != nil {
		fmt.Println("error getting response from the server")
		os.Exit(1)
	}
	
	defer res.Body.Close()


	body , err := io.ReadAll(res.Body)

	var rss RSSFeed

	if err := xml.Unmarshal(body , &rss); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	

	return &rss , nil 

}