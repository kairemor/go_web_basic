package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-yaml/yaml"
)

func main() {
	type Config struct {
		Database struct {
			Host string `json:"host" yaml:"host"`
			Port string `json:"port" yaml:"port"`
		} `json:"database" yaml:"database"`
		User struct {
			Username string `json:"login" yaml:"login"`
			Password string `json:"password" yaml:"password"`
		} `json:"user" yaml:"user"`
	}

	var configjson Config
	var configyaml Config

	filejson, err := os.Open("config.json")
	fileYaml, err := os.Open("config.yaml")

	if err != nil {
		fmt.Println(err)
	}

	defer filejson.Close()
	defer fileYaml.Close()

	content, err := ioutil.ReadAll(fileYaml)

	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(content, &configyaml)

	jsonDecoder := json.NewDecoder(filejson)
	jsonDecoder.Decode(&configjson)

	fmt.Println(configjson)
	fmt.Println(configyaml)

}
