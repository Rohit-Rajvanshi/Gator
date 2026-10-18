package main

import (
	"database/sql"
	"fmt"
	"os"
	"github.com/Rohit-Rajvanshi/Gator/internal/config"
	"github.com/Rohit-Rajvanshi/Gator/internal/database"
	_ "github.com/lib/pq"
)


func main() {
	con := config.Read()
	var s state
	s.c = &con 

	db , err := sql.Open("postgres" , s.c.DbURL)
	if err != nil {
		fmt.Printf("error connecting to the database")
		os.Exit(1)
	}
	
	dbQueries := database.New(db)

	s.db = dbQueries
	

	cmds := commands{
		comm: map[string]func(*state, command) error{},
	}
	
	cmds.register("login" , handlerLogin)
	cmds.register("register" , handlerRegister)
	cmds.register("reset" , handlerReset)
	cmds.register("users" , handlerListUsers)
	cmds.register("agg" , handlerAggregator)
	cmds.register("addfeed" , middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds" , handlerListFeeds)
	cmds.register("follow" , middlewareLoggedIn(handlerFollow))
	cmds.register("following" , middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow" , middlewareLoggedIn(handlerFeedUnfollow))
	cmds.register("browse" , handlerBrowse)

	if len(os.Args) < 2 {
		fmt.Println("Need more arguments!")
		os.Exit(1)
	}

	cmd := command{
		name : os.Args[1],
		arg : os.Args[2:],
	}

	err = cmds.run(&s , cmd)
	if err != nil {
		fmt.Println(err)
	}
}


