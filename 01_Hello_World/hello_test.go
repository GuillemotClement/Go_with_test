package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Jean")
	want := "Hello Jean"

	if got != want {
		t.Errorf("get %q, want %q", got, want)
	}
}
