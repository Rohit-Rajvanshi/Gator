package main

import (
	"context"
	"fmt"
	"os"
	"time"
    "github.com/google/uuid"
	"github.com/Rohit-Rajvanshi/Gator/internal/database"
)


func handlerFollow(s *state , cmd command , user database.User) error {
	if len(cmd.arg) == 0{
		fmt.Println("The register handler expects an argument, the username")
		os.Exit(1)
	}

	url := cmd.arg[0]
	ctx := context.Background()
	uuid := uuid.New()
	userID , err := s.db.GetUserID(ctx , user.Name)
	if err != nil {
		fmt.Println("error fetching the userID")
		os.Exit(1)
	}

	feedID , err  := s.db.GetFeedID(ctx , url)
	if err != nil {
		fmt.Println("error fetching the feedID")
		os.Exit(1)
	}

	

	feedFollow , err := s.db.CreateFeedFollow(ctx ,database.CreateFeedFollowParams{
		ID : uuid,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: userID,
		FeedID : feedID ,
	})

	fmt.Printf("Feed Name: %v\nUser Name: %v\n" , feedFollow.Name , feedFollow.Name_2)

	return nil 

}