package main

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func ExampleGroup_justErrors() {
	var g errgroup.Group
	var urls = []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://www.ryanair.com",
		"https://wizzair.com/",
		"https://www.swiss.com/",
	}
	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			// Fetch the URL.
			fmt.Println(url)
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	}
}

func ExampleGroup_noGoRoutine() {
	//var g errgroup.Group
	var urls = []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://www.ryanair.com",
		"https://wizzair.com/",
		"https://www.swiss.com/",
	}
	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		url := url // https://golang.org/doc/faq#closures_and_goroutines

		// Fetch the URL.
		fmt.Println(url)
		resp, err := http.Get(url)
		if err == nil {
			resp.Body.Close()
		} else {
			fmt.Errorf(err.Error())
			return
		}

	}
	// Wait for all HTTP fetches to complete.
	// if err := g.Wait(); err == nil {
	// 	fmt.Println("Successfully fetched all URLs.")
	// }
}
