package main

import (
	"context"
	"fmt"
	"os"
)


func handlerListUsers (s *state , _ command) error {
	ctx := context.Background()

	nameArray , err := s.db.ListUsers(ctx) 
	if err != nil {
		fmt.Println("error fetching names of the users from the datbase")
		os.Exit(1)
	}


	for _ , elem := range nameArray {
		if elem == s.c.CurrentUserName {
			fmt.Printf("* %s (current)\n" , elem)
			continue 
		}
		fmt.Println("*" , elem)
	}

	return nil 
}

