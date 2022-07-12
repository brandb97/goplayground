package main

import "fmt"
import "testing"

func TestMyTest(t *testing.T) {
	fmt.Println("Hello, world\nHi")
	t.Fatalf("Nooooooooooooooooooooooooooo")
}
