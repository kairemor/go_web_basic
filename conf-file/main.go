package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type Config struct {
		Database struct {
			Host string `json:"host"`
			Port string `json:"port"`
		} `json:"database"`
		User struct {
			Username string `json:"login"`
			Password string `json:"password"`
		} `json:"user"`
	}

	var config Config

	file, err := os.Open("config.json")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	jsonDecoder := json.NewDecoder(file)
	jsonDecoder.Decode(&config)

	fmt.Println(config)

}
