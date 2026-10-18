package main

import (
	"context"
	"github.com/Rohit-Rajvanshi/Gator/internal/database"
	"fmt"
	"os"
)


func handlerFeedUnfollow (s *state , cmd command , user database.User) error {
	ctx := context.Background()

	if len(cmd.arg) == 0{
		fmt.Println("Please provide the url of the feed")
		os.Exit(1)
	}

	feed_id , err  := s.db.GetFeedID(ctx , cmd.arg[0])
	if err != nil {
		return err 
	}

	err = s.db.FeedUnfollow(ctx , database.FeedUnfollowParams{
		UserID : user.ID ,
		FeedID: feed_id,
	} )

	if err != nil {
		return err 
	}

	return nil 

}