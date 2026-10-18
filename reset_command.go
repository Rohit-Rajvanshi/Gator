package main

import (
	"context"
	"fmt"
	"os"
)



func handlerReset (s *state , cmd command) error {
	ctx := context.Background()

	err := s.db.ResetTable(ctx) 

	if err != nil {
		fmt.Println("Deletion of the records was unsuccessful")
		os.Exit(1)
	}

	fmt.Println("Deletion of the records was successful")

	return nil 
}

