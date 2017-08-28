package client

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/go-test/deep"
)

func TestParse(t *testing.T) {
	request, error := http.NewRequest(
		"POST",
		"http://foo.example.com",
		bytes.NewBufferString("http://foo.bar,http://baz.qux"),
	)

	if error != nil {
		t.Error("Error creating a new request")
	}
	expected := []string{"http://foo.bar", "http://baz.qux"}
	parsed := Parse(request)
	if diff := deep.Equal(expected, parsed); diff != nil {
		t.Error(diff)
	}
}

// Once I actually know how to set up testing interfaces in Go and this project
// uses dependency injection
// TODO: TestRetrieveAll

// TODO: TestRetrieve
