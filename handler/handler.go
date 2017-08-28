package handler

import (
	"net/http"

	"github.com/MosesMendoza/gotextserver/client"
	"github.com/MosesMendoza/gotextserver/text"
)

/* handler package is our high level entry point for all logic related to
parsing, serving, and responding to a request. */

// Handle takes a req/response pair and processes using our helper packages
func Handle(w http.ResponseWriter, req *http.Request) {
	urls := client.Parse(req)
	content := client.RetrieveAll(urls)

	deduplicated := text.Deduplicate(content)
	sansIntegers := text.RemoveIntegers(string(deduplicated))
	// write our deduplicated bytes to the response
	w.Write([]byte(sansIntegers))
}
