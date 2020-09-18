package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://kairemor.gq"

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	fmt.Println(body)
}
