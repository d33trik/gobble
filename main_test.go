package main

import "testing"

func TestCountWords(t *testing.T) {
	input := "one two three four five"

	got := countWords([]byte(input))
	want := 5

	if got != want {
		t.Fail()
	}
}
