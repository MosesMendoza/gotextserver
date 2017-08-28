package handler

import (
	"net/http"
	"testing"
)

// implement a stub of the ResponseWriter interface

type stubResponseWriter struct{}

func (writer stubResponseWriter) Write(input []byte) (int, error) {
	return (len(input)), nil
}

func (writer stubResponseWriter) Header() http.Header {
	header := new(http.Header)
	header.Set("Content-Type", "text/plain")
	return *header
}

func (writer stubResponseWriter) WriteHeader(int) {}

// End stubResponseWriter implementation

func TestHandle(t *testing.T) {

	// TODO: Implement test that Handle calls expected functions

}
