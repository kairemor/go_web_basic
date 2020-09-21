package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/go-yaml/yaml"
)

func main() {
	type Config struct {
		Database struct {
			Host string `json:"host" yaml:"host"`
			Port string `json:"port" yaml:"port"`
		} `json:"database" yaml:"database"`
		User struct {
			Username string `json:"login" yaml:"login" toml:"login"`
			Password string `json:"password" yaml:"password"`
		} `json:"user" yaml:"user"`
	}

	var configjson Config
	var configyaml Config
	var configtoml Config

	filejson, err := os.Open("config.json")
	fileYaml, err := os.Open("config.yaml")
	fileToml, err := os.Open("config.toml")

	if err != nil {
		fmt.Println(err)
	}

	defer filejson.Close()
	defer fileYaml.Close()
	defer fileToml.Close()

	content, err := ioutil.ReadAll(fileYaml)
	content2, err := ioutil.ReadAll(fileToml)

	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(content, &configyaml)
	err = toml.Unmarshal(content2, &configtoml)

	jsonDecoder := json.NewDecoder(filejson)
	jsonDecoder.Decode(&configjson)

	fmt.Println(configjson)
	fmt.Println(configyaml)
	fmt.Println(configtoml)

}
