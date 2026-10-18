package main

import (
	"context"
	"html"
	"github.com/google/uuid"
	"github.com/Rohit-Rajvanshi/Gator/internal/database"
	"github.com/Rohit-Rajvanshi/Gator/internal/rss"
	"time"
	"database/sql"
	"strings"
	"log"
)


func scrapeFeeds(s *state) error {
	ctx := context.Background()

	nextFeed , err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return err 
	}

	err = s.db.MarkFeedFetched(ctx , nextFeed.ID)
	if err!= nil {
		return err 
	}

	fetchedFeed , err := rss.FetchFeed(ctx , nextFeed.Url)
	if err != nil {
		return err 
	}


	

	fetchedFeed.Channel.Title = html.UnescapeString(fetchedFeed.Channel.Title)
	fetchedFeed.Channel.Description = html.UnescapeString(fetchedFeed.Channel.Description)

	for i  := range fetchedFeed.Channel.Item{
		fetchedFeed.Channel.Item[i].Title = html.UnescapeString(fetchedFeed.Channel.Item[i].Title)
		fetchedFeed.Channel.Item[i].Description = html.UnescapeString(fetchedFeed.Channel.Item[i].Description)
	}



	
	for _ , elem := range fetchedFeed.Channel.Item {
		uuid := uuid.New()
		title := sql.NullString{
		String : elem.Title ,
		Valid : true ,
		}
		if title.String == "" {
			title.Valid = false
		}

		description := sql.NullString{
			String : elem.Description,
			Valid : true ,
		}
		if description.String == ""{
			description.Valid = false
		}

		publishedAt := sql.NullTime{}
		pubTime , err := time.Parse(time.RFC1123Z , elem.PubDate)
		if err != nil {
			publishedAt.Valid = false 
		}else{
			publishedAt.Time = pubTime
			publishedAt.Valid = true 
		}



		_ , err  = s.db.CreatePost(ctx , database.CreatePostParams{
			ID : uuid ,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: title ,
			Url:  elem.Link,
			Description: description,
			PublishedAt: publishedAt ,
			FeedID: nextFeed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error() , "duplicate key value violates unique constraint"){
				continue 
			}
			log.Print(err)
			continue 
		}
	}

	log.Printf("Feed %s collected, %v posts found", fetchedFeed.Channel.Title , len(fetchedFeed.Channel.Item))


	return nil 
}
