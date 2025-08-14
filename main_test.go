package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	got := hello_world()
	want := "Hello, World!"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	} else {
		t.Log("✅ Test passed successfully")
	}
}
