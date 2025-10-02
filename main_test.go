package main

import "testing"

func TestCountWords(t *testing.T) {
	input := "one two three four five"
	got := CountWords([]byte(input))
	want := 5

	if got != want {
		t.Logf("got: %d want: %d", got, want)
		t.Fail()
	}

	input = ""
	got = CountWords([]byte(input))
	want = 0

	if got != want {
		t.Logf("got: %d want: %d", got, want)
		t.Fail()
	}
}
