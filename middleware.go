package main

import (
	"context"
	"github.com/Rohit-Rajvanshi/Gator/internal/database"
)




func middlewareLoggedIn(handler func(s *state , cmd command , user database.User) error) func(*state , command) error {
	ctx := context.Background()

	return func(s *state , cmd command) error {
		user , err  := s.db.GetUser(ctx , s.c.CurrentUserName)
		if err != nil {
			return err 
		}
		err = handler(s , cmd , user)
		if err != nil{
			return err 
		}
		return nil 
	}
}