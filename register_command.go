package main

import (
	"context"
	"fmt"
	"os"
	"github.com/google/uuid"
	"time"
	"github.com/Rohit-Rajvanshi/Gator/internal/database"
)

func handlerRegister(s *state , cmd command) error {
	if len(cmd.arg) == 0{
		fmt.Println("The register handler expects an argument, the username")
		os.Exit(1)
	}
	ctx := context.Background()
	uuid := uuid.New()
	name := cmd.arg[0]

	if _ , err := s.db.GetUser(ctx , name); err == nil {
		fmt.Println("user already exists")
		os.Exit(1)
	}

	newUser , err := s.db.CreateUser(ctx , database.CreateUserParams{
		ID : uuid,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
	}) 
	if err != nil {
		fmt.Println("error creating the new user")
		os.Exit(1)
	}

	s.c.SetUser(name)

	fmt.Printf("user was created: %v\n" , newUser )
	
	
	return nil 
}