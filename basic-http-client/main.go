package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

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

func fetchWithChannel(urls []string, ch chan string) {
	for _, url := range urls {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Println(err)
			return
		}

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return
		}

		res := fmt.Sprintf("%10d", len(body))
		ch <- res
	}
	close(ch)
	// il est conseiller de fermer le channel que de la partie qui publie sur le channel
}

func main() {
	urls := []string{"http://kairemor.gq", "http://google.fr", "http://facebook.com"}

	// size := len(urls)
	// wg.Add(size)
	// for _, url := range urls {
	// 	go fetch(url)
	// }

	ch := make(chan string)
	go fetchWithChannel(urls, ch)

	// for

	// for l := range ch {
	// 	fmt.Println(l)
	// }

	ok := true
	for ok {
		select {
		case s, open := <-ch:
			{
				if !open {
					ok = false
				}
				fmt.Println(s)
			}
		}
	}
}
