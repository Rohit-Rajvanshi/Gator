package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/Rohit-Rajvanshi/Gator/internal/database"
)



func handlerBrowse(s *state , cmd command) error {
	ctx := context.Background()

	defaultLimit := ""

	if len(cmd.arg) == 0 {
		defaultLimit = "2"
	}else{
		defaultLimit = cmd.arg[0]
	}

	limit , err := strconv.ParseInt(defaultLimit , 10 , 64)
	if err!=nil {
		log.Printf("Please provide a valid limit: %s" , err )
	}

	userID , err  := s.db.GetUserID(ctx , s.c.CurrentUserName)
	if err != nil {
		log.Printf("Error fetching the user ID: %s" , err )
	}


	getPosts , err := s.db.GetPostsForUser(ctx , database.GetPostsForUserParams{
		UserID: userID,
		Limit: int32(limit),
	})
	if err != nil {
		log.Printf("Error fetching the posts for the user: %s" , err )
	}

	for _ , elem := range getPosts {
		fmt.Printf("\nPost:\nTitle: %v\nPublished Date: %v\nURL: %v\nDescription: %v\n", elem.Title.String , elem.PublishedAt.Time , elem.Url , elem.Description.String )
	} 

	return nil 
}


