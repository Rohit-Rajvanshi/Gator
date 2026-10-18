package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Rohit-Rajvanshi/Gator/internal/database"
)

func handlerFollowing(s *state, _ command , user database.User) error {

	ctx := context.Background()
	userID , err := s.db.GetUserID(ctx , user.Name) 
	if err != nil {
		fmt.Println("Error fetching the user_id")
		os.Exit(1)
	}

	followingList , err := s.db.GetFeedFollowsForUser(ctx , userID)
	if err != nil{
		fmt.Println("Error printing all the feed follows")
		os.Exit(1)
	}
	
	fmt.Println("User:" , user.Name)
	fmt.Println("Feeds followed:")

	for _ , elem := range followingList {
		fmt.Printf("- %v\n" , elem.Name)
	}

	return nil 
}

