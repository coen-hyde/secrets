package main

import (
	"testing"
)

func TestVersion(t *testing.T) {
	if Version != "dev" {
		t.Fatal("Version should be \"dev\"")
	}
}
