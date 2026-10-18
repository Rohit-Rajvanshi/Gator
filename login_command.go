package main

import (
	"fmt"
	"os"
	"context"
)


func handlerLogin (s *state , cmd command) error {
	if len(cmd.arg) == 0{
		fmt.Println("The login handler excepts a single argument, the username")
		os.Exit(1)
	}

	ctx := context.Background()

	if _ , err := s.db.GetUser(ctx , cmd.arg[0]); err != nil {
		fmt.Println("User is not registered")
		os.Exit(1)
	}
	


	s.c.SetUser(cmd.arg[0])

	fmt.Println("The user has been set!")

	return nil 
}



