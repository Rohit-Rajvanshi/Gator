# Gator - The Blog Aggregator 
Really fun challenge building this project as a part of the Bootdev Assignment 

## Requirements 
- Go 1.27.1
- PostgreSQL (running locally)

## Installation 
- `go install github.com/Rohit-Rajvanshi/Gator@latest`

## Setup 
- Create a Config file in your root directory `~/.gatorconfig.json`
```json 
{
  "db_url": "connection_string_goes_here",
  "current_user_name": "username_goes_here"
}
```
- It contains a JSON which has the database URL and the current User 
- Start the Postgres Server `sudo service postgresql start`
- Conect to the server (Use psql) `sudo -u postgres psql`
- Create a gator database `CREATE DATABASE gator;`
- Connect to the database `\c gator`
- Set the Password for the user(Linux/WSL only) `ALTER USER postgres PASSWORD 'postgres';`
- Run the migrations from the Gator/sql/schema directory UP : `goose postgres <connection_string> up` DOWN : `goose postgres <connection_string> down`
- Set the SSL mode to Disable `protocol://username:password@host:port/database?sslmode=disable` in the db_url in the JSON file




## Usage 
- `Gator register <name>` (As the name suggests , to register a user)
- `Gator login <user>` (To set the current user to the argument provided , if it was previously registered)
- ```Gator users``` (To show the list of all users)
- `Gator addfeed <name> <url>` (To add a feed to a User's Following Feeds , provide a custom name and the url to the feed)
- `Gator feeds` (To show the list of all the added feeds)
- `Gator follow <url>` (To follow a Feed that has already been added by a different user)
- `Gator following` (To show the feeds that user is following)
- `Gator unfollow <url>` (To unfollow a Feed that was previously added / followed by the current user)
- `Gator agg <interval>` (Aggregate the feeds added and create posts which gets added to the DB for later usage , It runs infinitely , use Ctrl+C to stop)(DO NOT USE TINY INTERVALS MIN: 1 Minute)
- `Gator browse <limit>` (Browse and List down the posts from the aggregated feeds for the current user , provide a limit to filter out the number of Posts displayed , default is 2)
- `Gator reset` (Delete everything)

## What I Learned
- Integrating a Go application with a PostgreSQL Database
- Using Sqlc and Goose 
- CTEs
- Middleware
- RSS / XML parsing


## Output 
![browse ouput](image.png)

