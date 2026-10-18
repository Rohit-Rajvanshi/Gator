package main 

import (
	"github.com/Rohit-Rajvanshi/Gator/internal/config"
	"github.com/Rohit-Rajvanshi/Gator/internal/database"
	"fmt"
)


type state struct {
	c *config.Config
	db *database.Queries
}

type command struct {
	name string 
	arg []string
}

type commands struct {
	comm map[string]func(*state , command) error 
}

func (c *commands) run(s *state, cmd command) error{
	val , ok := c.comm[cmd.name] 
	if ok == false {
		return fmt.Errorf("No such command exists")
	}
	return val(s , cmd)
}


func (c *commands) register(name string, f func(*state, command) error) {
	c.comm[name] = f 
}