package text

import "testing"

func TestDeduplicate(t *testing.T) {
	t.Log("Testing that basic text is deduplicated")
	var input = "abcabc123acb**4!"

	var deduplicated = Deduplicate(input)
	if !(deduplicated == "abc123*4!") {
		t.Error("Expected 'abc123*4!' got " + deduplicated)
	}
}

func TestEmptyInputToDeduplicate(t *testing.T) {
	t.Log("Testing that an empty string input results in an empty return")
	var deduplicated = Deduplicate("")
	if !(deduplicated == "") {
		t.Error("Expected empty string response, got " + deduplicated)
	}
}

func TestRemoveIntegers(t *testing.T) {
	t.Log("Testing removal of integers from a string")
	var integersRemoved = RemoveIntegers("abc123def345ghii")
	if !(integersRemoved == "abcdefghii") {
		t.Error("Expected abcdef but got " + integersRemoved)
	}
}
