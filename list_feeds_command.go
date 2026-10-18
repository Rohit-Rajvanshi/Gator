package main

import (
	"context"
	"fmt"
	"os"
)


func handlerListFeeds (s *state , _ command) error {
	ctx := context.Background()

	feedStruct , err := s.db.ListFeeds(ctx) 
	if err != nil {
		fmt.Println("error fetching feeds from the datbase")
		os.Exit(1)
	}


	for _ , elem := range feedStruct {
		fmt.Printf("FEED:\nname: %v\nurl: %v\nuser: %v\n" , elem.Name , elem.Url , elem.Name_2)
	}

	return nil 
}
