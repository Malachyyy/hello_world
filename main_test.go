package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	got := hello_world()
	want := "Hello, World!"
	fmt.Println("Test successful, got:", got)
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
