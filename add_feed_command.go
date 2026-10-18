package main

import (
	"context"
	"fmt"
	"os"
	"github.com/google/uuid"
	"github.com/Rohit-Rajvanshi/Gator/internal/database"
	"time"
)


func handlerAddFeed(s *state , cmd command , user database.User ) error {
	if len(cmd.arg) < 2{
		fmt.Println("The addFeed handler excepts two arguments, the name of the feed and the url")
		os.Exit(1)
	}



	ctx := context.Background()
	user_id , err := s.db.GetUserID(ctx , user.Name)
	if err != nil {
		fmt.Println("error getting the user_id")
		os.Exit(1)
	}

	name := cmd.arg[0]
	url := cmd.arg[1]
	uuid := uuid.New()

	newFeed , err := s.db.CreateFeed(ctx , database.CreateFeedParams{
		ID : uuid,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
		Url : url,
		UserID : user_id,
	})
	if err != nil {
		fmt.Println("error creating the new feed")
		os.Exit(1)
	}

	_ , feedErr := s.db.CreateFeedFollow(ctx ,database.CreateFeedFollowParams{
		ID : uuid,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user_id,
		FeedID : newFeed.ID ,
	})
	if feedErr != nil {
		fmt.Println("error following the feed")
		os.Exit(1)
	}

	fmt.Println(newFeed)

	return nil
}	