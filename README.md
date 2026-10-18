# Gator - The Blog Aggregator 

### Really fun challenge building this project as a part of the Bootdev Assignment 

## Requirements 
- Go 1.27.1
- PostgreSQL (running locally)

## Installation 
- go install https://github.com/Rohit-Rajvanshi/gator-blog-aggregator.git

## Setup 
- Create a Config file in your root directory ~/.gatorconfig.json
- It contains a JSON which has the database URL and the current User 
- Create a gator database 
- Run the migrations from the Gator/sql/schema directory 

## Usage 
- ```Gator register <gtname>``` (As the name suggests , to register a user)
- ```Gator login <user>``` (To set the current user to the argument provided , if it was previously registered)
- ```Gator users``` (To show the list of all users)
- ```Gator addfeed <name> <url>``` (To add a feed to a User's Following Feeds , provide a custom name and the url to the feed)
- ```Gator feeds``` (To show the list of all the added feeds)
- ```Gator follow <url>``` (To follow a Feed that has already been added by a different user)
- ```Gator unfollow <url>``` (To unfollow a Feed that was previously addded / followed by the current user)
- ```Gator agg <interval>``` (Aggregate the feeds added and create posts which gets added to the DB for later usage , It runs infinitely , use Ctrl+C to stop)(DO NOT USE TINY INTERVALS MIN: 1 Minute)
- ```Gator browse <limit>``` (Browse and List down the posts from the aggregated feeds for the current user , provide a limit to filter out the number of Posts displayed)
- ```Gator reset``` (Reset all the records)

## What i Learned / built with
- Integrating a Go application with a PostgreSQL Database
- Using Sqlc and Goose 

## Output 
![alt text](image.png)

