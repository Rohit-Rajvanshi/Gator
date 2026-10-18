package config 

import (
	"os"
	"log"
	"encoding/json"
)


type Config struct{
	DbURL string `json:"db_url"` 
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = "/.gatorconfig.json"


func Read() Config {
	
	file_path := getConfigFilePath()

	content , err := os.ReadFile(file_path)
	if err != nil {
		log.Fatalf("Failed to read the file: %v" , err)
	}

	var con Config 
	
	if err := json.Unmarshal(content , &con) ; err != nil {
		log.Fatalf("Error loading data into struct: %v" , err)
	}

	return con 
}

func (c *Config) SetUser(name string) {
	file_path := getConfigFilePath()
	c.CurrentUserName = name 
	jsonData , err := json.Marshal(c)
	if err != nil {
		log.Fatalf("Error marshaling data into json: %v" , err )
	}
	err = os.WriteFile(file_path , jsonData , 0644)
}

func getConfigFilePath() string {
	homeDir , err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to detect home directory: %v", err)
	}

	file_path := homeDir + configFileName

	return file_path
}




