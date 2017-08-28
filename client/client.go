package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

/* Our application needs to obtain the text content from one or more urls.
Effectively our server has its own client that obtains data referenced by an
incoming request and returns it for processing. That logic is captured here. */

// Parse takes a request and parses it into a []string of URL strings
func Parse(req *http.Request) []string {
	// Ensure we always close the request body even if we error
	defer req.Body.Close()

	// this is dangerous - we should be limiting the size of request we will read
	// into memory to prevent letting a giant request consume all our resources.
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		error := "Failed to read body from http request, error: " + err.Error()
		panic(error)
	}

	input := string(body)
	return strings.Split(input, ",")
}

// RetrieveAll takes a []string of URL strings and calls Retrieve on all of
// them, aggregating the associated responses.
func RetrieveAll(urls []string) string {
	responses := make(chan []byte, len(urls))
	var results []byte
	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go Retrieve(url, responses, &wg)
	}

	wg.Wait() // block until all content is retrieved
	close(responses)
	// retrieve all of the indivdual bodies from the channel, and then concatenate
	// them into a string
	for result := range responses {
		results = append(results, result...)

	}
	return string(results)
}

// Retrieve takes a url and retrieves the content, adding it to the results
// channel
func Retrieve(url string, results chan []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Processing request to retrieve content for url " + url)
	client := new(http.Client)
	resp, err := client.Get(url)
	if err != nil {
		panic("Could not retrieve content from url " + url)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("Could not read response body from url " + url)
	}
	results <- body
}
