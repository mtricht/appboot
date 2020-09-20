package main

import (
	"os"
	"testing"
)

func TestIsLauncher(t *testing.T) {
	os.Remove("appboot.json")
	answer := isLauncher()
	if answer {
		t.Errorf("isLauncher() = %v; want false", answer)
	}
	file, _ := os.Create("appboot.json")
	file.Close()
	answer = isLauncher()
	if !answer {
		t.Errorf("isLauncher() = %v; want true", answer)
	}
	os.Remove("appboot.json")
}
