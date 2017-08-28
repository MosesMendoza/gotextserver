package main

/* main exists only to set up our http server and begin processing requests. It
delegates all logic to the handler. */
import (
	"net/http"

	"github.com/MosesMendoza/gotextserver/handler"
)

func main() {
	http.Handle("/", http.HandlerFunc(handle))
	http.ListenAndServe("0.0.0.0:8000", nil)
}

// handle hands off the request to our handler for processing
func handle(w http.ResponseWriter, req *http.Request) {
	handler.Handle(w, req)
}
