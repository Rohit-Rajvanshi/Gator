package main

import (
	"fmt"
	"time"	
	"os"
)



func handlerAggregator (s *state , cmd command ) error {
	if len(cmd.arg) == 0 {
		fmt.Println("The agg handler expects an argument, the time duration")
		os.Exit(1)
	}

	timeBetweenReqs , err := time.ParseDuration(cmd.arg[0])
	if err != nil {
		return err 
	}


	fmt.Printf("Collecting feeds every , %s\n" , cmd.arg[0])

	ticker := time.NewTicker(timeBetweenReqs)

	for ;; <- ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			fmt.Println(err)
			continue 
		}
	}
}


