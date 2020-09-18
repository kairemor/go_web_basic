package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	urls := []string{"http://kairemor.gq", "http://google.fr", "http://facebook.com"}

	size := len(urls)
	wg.Add(size)
	for _, url := range urls {
		go fetch(url)
	}

}

func fetch(url string) {
	defer wg.Done()
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
